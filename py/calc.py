import re

codehelp = [
    "This Code Help:",
    "加法: +",
    "減法: -",
    "乗法: *",
    "除法: /",
    "累乗: ^2, 2のところはそれぞれ置換してください",
    "パーセント: +10%, 数字、演算子はそれぞれ置換してください",
    "最初の数値: "
]

def add(x, y):
    return x + y

def subtract(x, y):
    return x - y

def multiply(x, y):
    return x * y

def divide(x, y):
    if y != 0:
        return x / y
    else:
        return "Error! Division by zero."

def power(x, exp):
    return x ** exp

def percent_increase(x, percent):
    return x * (1 + percent / 100)

def format_number(num):
    return "{:,}".format(num)

def clean_number(number_str):
    return float(number_str.replace(',', ''))

def calculator():
    # 演算子と対応する関数の辞書
    enzankigo = {
        '+': add,
        '-': subtract,
        '*': multiply,
        '/': divide
    }
    
    # 特殊な操作の辞書
    special_operations = {
        '^': power,
        '%': percent_increase
    }
    
    # 最初の数値を取得
    num1 = clean_number(input("\n".join(codehelp)))
    
    while True:
        print("演算子と数値を入力してください（例: +111、-222、*300、/400、^2、+60%）。終了するには 'n' を入力してください。")
        user_input = input("入力: ").strip()
        
        if user_input.lower() == 'n':
            break
        
        # 累乗のチェック
        match = re.match(r"(\d+)(?:\^(\d*))?", user_input)
        if match:
            base = clean_number(match.group(1))
            exponent_str = match.group(2)
            exponent = int(exponent_str) if exponent_str else 2  # デフォルトの指数を2に設定
            result = power(base, exponent)
            print(f"Output: {format_number(result)}")
            num1 = result
            continue
        
        # パーセンテージ増加のチェック
        if user_input.endswith('%'):
            try:
                percent = float(user_input[:-1])
                result = percent_increase(num1, percent)
                print(f"Output: {format_number(result)}")
                num1 = result
            except ValueError:
                print("パーセンテージの入力が無効です。")
            continue
        
        # 演算子と数値の入力
        if user_input:
            operator = user_input[0]
            try:
                number = clean_number(user_input[1:])
            except ValueError:
                print("数値の入力が無効です。")
                continue
            
            if operator in enzankigo:
                result = enzankigo[operator](num1, number)
                if isinstance(result, str):
                    print(result)  # エラーメッセージを表示
                else:
                    print(f"Output: {format_number(result)}")
                    num1 = result  # 現在の結果を次の計算の入力として使用
            else:
                print("無効な演算子です。")
        else:
            # 入力が空の場合はデフォルトで加法を適用
            number = float(input("数値を入力してください: "))
            result = add(num1, number)
            print(f"Output: {format_number(result)}")
            num1 = result  # 現在の結果を次の計算の入力として使用

    # 最終結果の表示
    print(f"最終計算結果: {format_number(num1)}")

# 電卓を起動
calculator()