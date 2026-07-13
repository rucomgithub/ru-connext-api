# Deploy Notes

## ⚠️ `docker-compose.yml` ที่ commit ไว้เป็น dev-mode — ไม่ auto-deploy จาก git

```yaml
volumes:
   - ./:/app                # bind-mount source code จาก host ตรงๆ
   - ./logger:/logger
command: reflex -r "\.go$$" -s -- sh -c "go run ./"
```

Container รันโค้ดจากโฟลเดอร์บน host ตรงๆ ผ่าน `go run` (ไม่ใช่ image ที่ build ไว้ล่วงหน้า) และ `reflex` จะ compile+restart ให้เองอัตโนมัติ **เฉพาะเมื่อไฟล์ `.go` เปลี่ยน** (pattern `-r "\.go$"`)

**ผลที่ตามมาที่พลาดกันบ่อย:**

1. **`git push`/merge PR ไม่ทำให้ server รันโค้ดใหม่เอง** — ต้อง `git pull` บนโฟลเดอร์ที่ mount เข้า container ด้วยตัวเอง (ไม่มี CI/CD auto-deploy ผูกกับ repo นี้)
2. **แก้ `environments/config.yaml` เฉยๆ ไม่พอ** — `reflex` เฝ้าดูแค่ไฟล์ `.go` ไม่รวม `.yaml` และ viper อ่าน config แค่ครั้งเดียวตอน process เริ่ม (`environments.EnvironmentInit()`, ไม่มี `viper.WatchConfig()`) → แก้ config แล้วต้อง **restart container เอง** ถึงจะมีผล แม้ไฟล์บน host จะเปลี่ยนทันทีเพราะเป็น bind mount

**Checklist หลังแก้โค้ดหรือ config:**

```bash
cd <โฟลเดอร์ที่ docker-compose.yml อยู่บน server>
git pull                                    # ดึงโค้ดใหม่ (ถ้าแก้โค้ด)
# แก้ environments/config.yaml ตามต้องการ (ถ้ามี)
docker compose restart ru-smart-api         # บังคับ restart เสมอ — ทั้งสองกรณี
```

ยืนยันว่า container รันโค้ด/ค่าที่ต้องการแล้วจริง:
```bash
docker compose exec ru-smart-api git log -1 --oneline    # เทียบ commit กับที่ push ไป
docker compose exec ru-smart-api cat environments/config.yaml | grep <key>
```

## Google Sign-In: มี config key ชื่อคล้ายกันสองตัว

`environments/config.yaml` มี `google.google_client_id` และเคยมี `google.google_client_appid` (ถูกลบออกจากโค้ดแล้ว — ดู `middlewares/googleAuth.go`) — **ใช้ `google.google_client_id` เพียงตัวเดียว** เป็น audience (`aud`) ที่ backend เช็คตอน validate Google ID token ที่ `/google/authorization`

ค่าที่ถูกต้องคือ **Web application OAuth client (type 3)** ของ Google Cloud project เดียวกับที่แอป mobile ใช้สร้าง Firebase/`google-services.json` (ปัจจุบันคือ project `630106548850`) — ไม่ใช่ Android/iOS client ID เพราะ `aud` ของ ID token อ้างอิงจาก Web client เสมอไม่ว่าผู้ใช้ login จาก Android หรือ iOS
