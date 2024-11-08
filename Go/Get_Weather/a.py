import re
import json

def parse_text_to_json(filename):
    data = {}

    with open(filename, 'r', encoding='utf-8') as file:
        text_data = file.read()
    
    # 各 <pref> ブロックを取得
    pref_blocks = re.findall(r'<pref title="([^"]+)">(.*?)</pref>', text_data, re.DOTALL)
    
    for pref_title, pref_content in pref_blocks:
        data[pref_title] = {}
        
        # <city> タグの title と id を抽出
        city_matches = re.findall(r'<city title="([^"]+)" id="([^"]+)"', pref_content)
        for city_title, city_id in city_matches:
            data[pref_title][city_title] = city_id

    # JSONとして出力
    json_output = json.dumps(data, ensure_ascii=False, indent=2)
    print("結果のJSON出力:")
    print(json_output)

# ファイルから解析
parse_text_to_json('a.txt')
