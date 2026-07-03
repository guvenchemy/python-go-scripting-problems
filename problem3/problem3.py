import json
from collections import Counter

success_counter = 0
error_counter = 0
level_counter, status_counter = Counter(), Counter()
error_paths = set()
total_response_time = 0

with open("app_logs.json", "r", encoding="utf-8") as file:
        for line in file:
            line = line.strip()
            if line:
                try:
                    item = json.loads(line)
                    level_counter[item["level"]]+=1
                    status_counter[item["status"]]+=1
                    if item["level"] == "ERROR":
                        error_paths.add(item["path"])
                    total_response_time += item["response_time_ms"]
                    success_counter += 1
                except json.JSONDecodeError as e:
                    print(f"Error: {e}")
                    error_counter += 1
                    pass


print("Log seviyeleri:", level_counter)
print("Status kodlari:",status_counter)
print("error_paths:",error_paths)
print("Başarıyla işlenen log satır sayısı:", success_counter)
print("Hatalı log satır sayısı:", error_counter)
if success_counter > 0:
    print("Ortalama Yanıt Süresi (ms):", total_response_time / success_counter)
else:
    print("Ortalama Yanıt Süresi (ms): 0 (Başarılı log bulunamadı)")