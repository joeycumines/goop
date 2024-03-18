/*
make_lib.go
Description:
	An implementation of the file make_lib.go written entirely in go.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MakeLibFlags struct {
	GoFilename  string // Name of the Go File to
	PackageName string // Name of the package
}

type LPSolveVersionInfo struct {
	MajorVersion int
	MinorVersion int
}

func GetDefaultMakeLibFlags() (MakeLibFlags, error) {
	// Create Default Struct
	mlf := MakeLibFlags{
		GoFilename:  "solvers/lib.go",
		PackageName: "solvers",
	}

	return mlf, nil
}

func ParseMakeLibArguments(mlfIn MakeLibFlags) (MakeLibFlags, error) {
	// Iterate through any arguments with mlfIn as the default
	mlfOut := mlfIn

	// Input Processing
	argIndex := 1 // Skip entry 0
	for argIndex < len(os.Args) {
		// Share parsing data
		fmt.Printf("- Parsed input: %v\n", os.Args[argIndex])

		// Parse Inputs
		switch {
		case os.Args[argIndex] == "--go-fname":
			mlfOut.GoFilename = os.Args[argIndex+1]
			argIndex += 2
		case os.Args[argIndex] == "--pkg":
			mlfOut.PackageName = os.Args[argIndex+1]
			argIndex += 2
		default:
			fmt.Printf("Unrecognized input: %v", os.Args[argIndex])
			argIndex++
		}
	}

	return mlfOut, nil
}

/*
CreateCXXFlagsDirective
Description:

	Creates the CXX Flags directive in the  file that we will use in lib.go.
*/
func CreateCXXFlagsDirective(mlfIn MakeLibFlags) (string, error) {
	// Create Statement
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	pwdCXXFlagsString := fmt.Sprintf("// #cgo CXXFLAGS: --std=c++11 -I%v/include\n", pwd)
	lpSolveCXXFlagsString := "// #cgo CXXFLAGS: -I/opt/homebrew/opt/lp_solve/include\n" // Works as long as lp_solve was installed with Homebrew

	return fmt.Sprintf("%v%v", pwdCXXFlagsString, lpSolveCXXFlagsString), nil
}

/*
CreatePackageLine
Description:

	Creates the "package" directive in the  file that we will use in lib.go.
*/
func CreatePackageLine(mlfIn MakeLibFlags) (string, error) {
	return fmt.Sprintf("package %v\n\n", mlfIn.PackageName), nil
}

/*
CreateLDFlagsDirective
Description:

	Creates the LD_FLAGS directive in the file that we will use in lib.go.
*/
func CreateLDFlagsDirective(mlfIn MakeLibFlags) (string, error) {
	var ldFlagsDirective string
	// Write the lp_solve LD Flags line.
	lvi, err := DetectLPSolveVersion()
	if err != nil {
		return "", err
	}
	ldFlagsDirective = fmt.Sprintf("%v// #cgo LDFLAGS: -L/opt/homebrew/opt/lp_solve/lib -llpsolve%v%v\n", ldFlagsDirective, lvi.MajorVersion, lvi.MinorVersion)
	return ldFlagsDirective, nil
}

/*
HeaderNameToLPSolveVersionInfo
Description:

	Converts the header file (like liblpsolve55.a) into an LPSolveVersionInfo object which can be used later.
*/
func HeaderNameToLPSolveVersionInfo(lpsolveHeaderName string) (LPSolveVersionInfo, error) {
	//Locate major and minor version indices in directory name
	majorVersionAsString := string(lpsolveHeaderName[len("liblpsolve")])
	minorVersionAsString := string(lpsolveHeaderName[len("liblpsolve")+1])

	// Convert using strconv to integers
	majorVersion, err := strconv.Atoi(majorVersionAsString)
	if err != nil {
		return LPSolveVersionInfo{}, err
	}

	minorVersion, err := strconv.Atoi(minorVersionAsString)
	if err != nil {
		return LPSolveVersionInfo{}, err
	}

	return LPSolveVersionInfo{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
	}, nil
}

func GetAHeaderFilenameFrom(dirName string) (string, error) {
	// Constants

	// Algorithm

	// Search through dirName directory for all instances of .a files
	libraryContents, err := os.ReadDir(dirName)
	if err != nil {
		return "", err
	}
	headerNames := []string{}
	for _, content := range libraryContents {
		if content.Type().IsRegular() && strings.Contains(content.Name(), ".a") {
			fmt.Println(content.Name())
			headerNames = append(headerNames, content.Name())
		}
	}

	return headerNames[0], nil

}

func DetectLPSolveVersion() (LPSolveVersionInfo, error) {
	// Constants
	homebrewLPSolveDirectory := "/opt/homebrew/opt/lp_solve"

	// Algorithm
	headerFilename, err := GetAHeaderFilenameFrom(fmt.Sprintf("%v/lib/", homebrewLPSolveDirectory))
	if err != nil {
		return LPSolveVersionInfo{}, err
	}

	return HeaderNameToLPSolveVersionInfo(headerFilename)

}

func WriteLibGo(mlfIn MakeLibFlags) error {
	// Constants

	// Algorithm

	// First Create all Strings that we would like to write to lib.go
	// 1. Create package definition
	packageDirective, err := CreatePackageLine(mlfIn)
	if err != nil {
		return err
	}

	// 2. Create CXX_FLAGS argument
	cxxDirective, err := CreateCXXFlagsDirective(mlfIn)
	if err != nil {
		return err
	}

	// 3. Create LDFLAGS Argument
	ldflagsDirective, err := CreateLDFlagsDirective(mlfIn)
	if err != nil {
		return err
	}

	// Now Write to File
	f, err := os.Create(mlfIn.GoFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write all directives to file
	_, err = f.WriteString(fmt.Sprintf("%v%v%vimport \"C\"\n", packageDirective, cxxDirective, ldflagsDirective))
	if err != nil {
		return err
	}

	return nil

}
func main() {
	mlf, err := GetDefaultMakeLibFlags()
	if err != nil {
		panic(err)
	}

	// Next, parse the arguments to make_lib and assign values to the mlf appropriately
	mlf, err = ParseMakeLibArguments(mlf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", mlf)

	// Write File
	err = WriteLibGo(mlf)
	if err != nil {
		panic(err)
	}
}
