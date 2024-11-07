import GPUtil

def get_gpu_status():
    gpus = GPUtil.getGPUs()
    for gpu in gpus:
        print(f"GPU Name: {gpu.name}")
        print(f"  Load: {gpu.load * 100}%")
        print(f"  Free Memory: {gpu.memoryFree}MB")
        print(f"  Used Memory: {gpu.memoryUsed}MB")
        print(f"  Total Memory: {gpu.memoryTotal}MB")
        print(f"  Temperature: {gpu.temperature} °C")

# GPUの状態を取得して表示
get_gpu_status()
