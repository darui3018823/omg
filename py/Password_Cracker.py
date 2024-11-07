import asyncio
import itertools
import string
import time
import os
import cupy as cp
from concurrent.futures import ProcessPoolExecutor

# 使える文字のセット（Unicode文字列）
charset = string.ascii_uppercase + string.ascii_lowercase + string.digits + "-."

# パスワードの組み合わせを生成してチェックする関数
def gpu_brute_force(target_password, length):
    target_gpu = cp.array(list(target_password.encode('utf-8')), dtype=cp.uint8)  # dtypeをcp.uint8に設定
    guesses = itertools.product(charset, repeat=length)

    for guess_tuple in guesses:
        guess_str = ''.join(guess_tuple)  # バイトに結合
        guesses_gpu = cp.array(list(guess_str.encode('utf-8')), dtype=cp.uint8)  # dtypeをcp.uint8に設定

        # GPUで比較
        if cp.array_equal(guesses_gpu, target_gpu):
            print(f"Password: {guess_str} Success!")
            return guess_str  # 完全一致した場合に返す

        # 各試行結果を出力
        print(f"Password: {guess_str} Attack Failure.")

    return None

# CPUとGPUの情報を表示
def display_system_info():
    print(f"CPU Core Count: {os.cpu_count()}")
    try:
        gpu = cp.cuda.Device(0)
        gpu_properties = cp.cuda.runtime.getDeviceProperties(0)
        gpu_name = gpu_properties['name'].decode()
        gpu_memory = gpu_properties['totalGlobalMem'] / (1024 ** 3)
        print(f"GPU: {gpu_name}")
        print(f"VRAM: {gpu_memory:.2f} GB")
    except Exception as e:
        print("Error retrieving GPU information:", e)

# 非同期で総当たり攻撃を実行
async def async_brute_force_password(target_password, mode):
    start_time = time.time()
    max_workers = os.cpu_count() * 2  # CPUコアの2倍のプロセスを使用

    with ProcessPoolExecutor(max_workers=max_workers) as executor:
        loop = asyncio.get_event_loop()

        if mode == "normal":
            # 通常モード: 1桁から順に試行
            for length in range(1, len(target_password) + 1):
                print(f"Trying length: {length}")
                future = loop.run_in_executor(executor, gpu_brute_force, target_password, length)
                result = await future
                if result:
                    break
        elif mode == "easy":
            # 倖田來未モード: 入力された桁数で試行
            length = len(target_password)
            print(f"Trying length: {length}")
            future = loop.run_in_executor(executor, gpu_brute_force, target_password, length)
            result = await future

        if result:
            elapsed_time = time.time() - start_time
            print(f"パスワードが見つかりました: {result}")
            print(f"かかった時間: {elapsed_time:.3f} 秒")
        else:
            print("パスワードが見つかりませんでした")

# メイン部分
if __name__ == "__main__":
    display_system_info()
    
    while True:
        target_password = input("ターゲットパスワードを入力してください (終了するには 'exit' と入力): ")
        if target_password.lower() == "exit":
            print("プログラムを終了します。")
            break

        mode = input("モードを選択してください (通常: 'normal', EASYモード: 'easy'): ")
        asyncio.run(async_brute_force_password(target_password, mode))
