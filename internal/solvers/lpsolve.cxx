#include "lp_lib.h"
#include "lpsolve.hpp"
#include <iostream>
#include <map>

using namespace std;

LPSolveSolver::LPSolveSolver() : lp(nullptr), numVars(0), logLevel(NEUTRAL) {}

LPSolveSolver::~LPSolveSolver()
{
    if (lp != nullptr)
    {
        delete_lp(lp);
    }
}

void LPSolveSolver::showLog(bool shouldShow)
{
    logLevel = shouldShow ? FULL : NEUTRAL;
}

void LPSolveSolver::setTimeLimit(double timeLimit)
{
    set_timeout(lp, timeLimit);
}

void LPSolveSolver::addVars(int count, double *lb, double *ub, char *types)
{
    if (lp != nullptr) {
        delete_lp(lp);
    }

    lp = make_lp(0, count);
    set_verbose(lp, logLevel);
    set_add_rowmode(lp, TRUE);
    numVars = count;

    for (size_t i = 0; i < count; i++)
    {
        set_bounds(lp, i + 1, lb[i], ub[i]);
        switch (types[i])
        {
            case 'I':
                set_int(lp, i + 1, TRUE);
                break;
            case 'B':
                set_binary(lp, i + 1, TRUE);
                break;
        }
    }
}

void LPSolveSolver::addConstr(
    int lhs_count, double* lhs_coeffs, uint64* lhs_var_ids, double lhs_constant,
    int rhs_count, double* rhs_coeffs, uint64* rhs_var_ids, double rhs_constant, char sense)
{
    std::map<int, double> net_coeffs;

    // Process LHS variables
    for (int i = 0; i < lhs_count; ++i) {
        int var_id = (int)lhs_var_ids[i] + 1; // Adjusting for 1-based indexing in LPSolve
        net_coeffs[var_id] += lhs_coeffs[i];
    }

    // Process RHS variables, subtracting their coefficients
    for (int i = 0; i < rhs_count; ++i) {
        int var_id = (int)rhs_var_ids[i] + 1; // Adjusting for 1-based indexing in LPSolve
        net_coeffs[var_id] -= rhs_coeffs[i];
    }

    // Prepare the arrays for add_constraintex call
    int var_count = net_coeffs.size();
    REAL sparse_row[var_count];
    int colno[var_count];
    int index = 0;

    for (const auto& kv : net_coeffs) {
        colno[index] = kv.first;
        sparse_row[index] = kv.second;
        ++index;
    }

    // Determine constraint type
    int constr_type;
    switch (sense) {
        case '=': constr_type = EQ; break;
        case '<': constr_type = LE; break;
        case '>': constr_type = GE; break;
        default: throw std::invalid_argument("Unsupported constraint sense.");
    }

    // Calculate the constant part of the constraint
    REAL constant = rhs_constant - lhs_constant;

    // Add the constraint to the model
    add_constraintex(lp, var_count, sparse_row, colno, constr_type, constant);
}

void LPSolveSolver::setObjective(
    int count, double *coeffs, uint64 *var_ids, double constant, int sense)
{
    REAL sparse_row[count];
    int colno[count];

    for (size_t i = 0; i < count; i++)
    {
        sparse_row[i] = coeffs[i];
        colno[i] = (int) var_ids[i] + 1;
    }

    set_obj_fnex(lp, count, sparse_row, colno);

    switch (sense)
    {
        case 1:
            set_minim(lp);
            break;
        case -1:
            set_maxim(lp);
            break;
    }
}

MIPSolution LPSolveSolver::optimize()
{
    set_add_rowmode(lp, false);

    if (logLevel == FULL) {
        write_lp(lp, nullptr);
    }

    int res = solve(lp);

    MIPSolution sol;
    sol.optimal = res == OPTIMAL;
    sol.gap = get_mip_gap(lp, TRUE);
    sol.errorCode = res;
    sol.errorMessage = "No error messages provided for LPSolve";

    sol.values.resize(numVars);
    REAL vars[numVars];
    get_variables(lp, vars);
    for (size_t i = 0; i < numVars; i++)
    {
        sol.values.at(i) = (double) vars[i];
    }

    return sol;
}
