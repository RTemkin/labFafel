from itertools import product

# Функция для проверки, является ли строка логической переменной
def is_variable(s):
    return s.isalpha() and len(s) == 1

# Функция для оценки логического выражения с заданным контекстом
def eval_expression(expr, context):
    # Заменяем переменные на их значения
    for var in context:
        expr = expr.replace(var, str(context[var]))
    # Выполняем вычисления
    expr = expr.replace('¬', ' not ')
    expr = expr.replace('∧', ' and ')
    expr = expr.replace('∨', ' or ')
    expr = expr.replace('→', ' <= ')  # Импликация: A → B = ¬A ∨ B
    expr = expr.replace('↔', ' == ')  # Эквиваленция: A ↔ B = (A → B) ∧ (B → A)
    return eval(expr)

# Функция для преобразования логического выражения в ДНФ
def to_dnf(expr):
    expr = expr.replace('AND', '∧').replace('OR', '∨').replace('NOT', '¬')
    variables = set(filter(is_variable, expr))
    table = []

    for combination in product([False, True], repeat=len(variables)):
        context = dict(zip(variables, combination))
        truth_value = eval_expression(expr, context)
        table.append((combination, truth_value))

    return variables, table

# Функция для преобразования логического выражения в КНФ
def to_knf(expr):
    expr = expr.replace('AND', '∧').replace('OR', '∨').replace('NOT', '¬')
    variables = set(filter(is_variable, expr))
    table = []

    for combination in product([False, True], repeat=len(variables)):
        context = dict(zip(variables, combination))
        truth_value = eval_expression(expr, context)
        table.append((combination, truth_value))

    return variables, table

# Функция для формирования ДНФ из таблицы истинности
def get_dnf(variables, table):
    dnf_terms = []
    for combination, value in table:
        if value:  # Извлекаем строки, где выражение истинно
            term = []
            for var, state in zip(variables, combination):
                if state:  # Если переменная истинна
                    term.append(var)
                else:  # Если переменная ложна
                    term.append(f'¬{var}')
            dnf_terms.append('(' + ' ∧ '.join(term) + ')')
    return ' ∨ '.join(dnf_terms)

# Функция для формирования КНФ из таблицы истинности
def get_knf(variables, table):
    knf_clauses = []
    for combination, value in table:
        if not value:  # Извлекаем строки, где выражение ложно
            clause = []
            for var, state in zip(variables, combination):
                if state:  # Если переменная истинна
                    clause.append(f'¬{var}')
                else:  # Если переменная ложна
                    clause.append(var)
            knf_clauses.append('(' + ' ∨ '.join(clause) + ')')
    return ' ∧ '.join(knf_clauses)

# Функция для формирования таблицы истинности
def print_truth_table(variables, table):
    header = ' | '.join(variables) + ' | Result'
    print(header)
    print('-' * len(header))
    for combination, truth_value in table:
        truth_row = ' | '.join(str(int(val)) for val in combination)
        print(f"{truth_row} | {int(truth_value)}")

def main():
    logical_expression = input("Введите логическое выражение (используйте AND, OR, NOT, →, ↔) например A ∧ (B ∨ ¬C): ")
    
    variables, truth_table = to_dnf(logical_expression)
    print("\nТаблица истинности:")
    print_truth_table(variables, truth_table)
    
    dnf_result = get_dnf(variables, truth_table)
    print("\nДизъюнктивная нормальная форма (ДНФ):")
    print(dnf_result)
    
    knf_result = get_knf(variables, truth_table)
    print("\nКонъюнктивная нормальная форма (КНФ):")
    print(knf_result)


if __name__ == "__main__":
    main()


