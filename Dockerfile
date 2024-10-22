# ใช้ image ของ PostgreSQL
FROM postgres:14

# ตั้งค่าตัวแปรแวดล้อมสำหรับการสร้างฐานข้อมูล เบื้องต้น
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase


# เปิด port ที่ PostgreSQL ใช้งาน
EXPOSE 5432
