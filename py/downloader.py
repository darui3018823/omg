import subprocess
import os

def download_video(url):
    # ダウンロード先ディレクトリを指定
    download_dir = 'E:/yt-dlp'

    # yt-dlp.exe のフルパスを指定
    ytdlp_path = 'yt-dlp'

    # URLに基づいてコマンドリストを変更
    if "twitch.tv" in url:
        command = [
            ytdlp_path,
            '-f', '1080p60+bestaudio',  # 最大画質を指定
            '--merge-output-format', 'mp4',
            '--embed-thumbnail',  # サムネイルを埋め込む
            '--add-metadata',  # メタデータを追加
            '--output', os.path.join(download_dir, '%(title)s.%(ext)s'),  # E:\yt-dlp\1 に保存
            url
        ]
    elif "youtube.com" in url or "youtu.be" in url:
        command = [
            ytdlp_path,
            '-f', 'bestvideo+bestaudio',
            '--merge-output-format', 'mp4',
            '--embed-thumbnail',  # サムネイルを埋め込む
            '--add-metadata',  # メタデータを追加
            '--output', os.path.join(download_dir, '%(title)s.%(ext)s'),  # E:\yt-dlp\1 に保存
            url
        ]
    elif "twitter.com" in url or "x.com" in url:
        command = [
            ytdlp_path,
            '--merge-output-format', 'mp4',
            '--embed-thumbnail',  # サムネイルを埋め込む
            '--add-metadata',  # メタデータを追加
            '--output', os.path.join(download_dir, '%(title)s.%(ext)s'),  # E:\yt-dlp\1 に保存
            url
        ]
    else:
        print("対応していないURLです。")
        return

    # コマンドの実行
    try:
        subprocess.run(command, check=True)
        print("ダウンロードが完了しました。")
    except subprocess.CalledProcessError as e:
        print(f"エラーが発生しました: {e}")

if __name__ == "__main__":
    while True:
        video_url = input("ダウンロードしたい動画のURLを入力してください (終了するには 'exit' と入力): ")
        
        if video_url.lower() in ['exit', 'q']:
            print("プログラムを終了します。")
            break
        
        download_video(video_url)
