# ตัวอย่างการใช้งาน Demo 💢
### 🔘 step 1
รัน docker contrainer ของ redis 
	
> * ` docker run --name some-redis -d redis`

### 🔘 step 2
> * `cd sub`
> * `go run subscriber.go`

### 🔘 step 3
> * `cd pub`
> * `go run publisher.go`

✅ ผลลัพธ์คือ subscriber จะได้รับ message ที่ publisher ทำการ publish ไป และการทำงานจะเป็นไปตามที่อธิบายด้านล่างนี้

## Messaging Broker by using REDIS 🔴
> *  ref : https://redis.io/docs/latest/develop/interact/pubsub/

การทำงานของ redis ในโหมดของ Messaging Broker คือ
1. Publisher ส่ง message ไปยัง redis
2. Redis Node จะเช็ค key และ client ที่ Subscribe
3. Redis จะส่ง message ไปยัง subscriber

![image](https://github.com/user-attachments/assets/c9ab41fc-d4d8-40f4-897b-7625450f4a4b)




## Comparing Redis and RabbitMQ as Messaging Brokers
##### เมื่อเปรียบเทียบระหว่าง Redis และ RabbitMQ เมื่อใช้โหมด messaging broker
> * ref: https://aws.amazon.com/compare/the-difference-between-rabbitmq-and-redis/

Feature | Redis | RabbitMQ
------------ | ------------- | -------------
Type(ประเภท) | เป็น In-memory data store ที่มีระบบ publish-subscribe (pub/sub) messaging system  | เป็น Message broker ที่ใช้ protocal AMQP 
Message Pattern | Pub/Sub with channel | Producer/Consumer with routing, exchanges, and queues
Performance(ประสิทธิภาพ) | เร็วมาก เพราะ low latency | มีการทำงานที่ซับซ้อนกว่าทำให้ช้ากว่า
Persistence(ความยั่งยืนของข้อมูล) | ข้อมูลสามารถหายได้ | ข้อมูลจะคงทนกว่า redis
Message Ordering(การจัดเรียงข้อมูล) | ไม่มีการจัดเรียงข้อมูลการทำงาน | มีการจัดเรียงข้อมูลแบบ Queue หรือ FIFO
Message Acknowledgment(การรับทราบข้อความ) | ไม่มีระบบการรับทราบข้อมูล(Ack) | มีระบบการรับทราบข้อมูล(Ack)
Reliability(ความน่าเชื่อถือ) | ไม่การันตีการรับส่งข้อมูล | มีความน่าเชื่อถือเพราะมีระบบการรับทราบข้อมูล(Ack) และระบบคิว(Queue)
Latency(ระยะเวลาในการส่งจนถึงรับข้อมูล) | <1ms (extremely fast) | ~10ms (slightly higher latency)
Throughput(ปริมาณการทำงาน) | High (1M messages per second) | High (100,000 messages per second)
Message size(ขนาดของข้อมูล) | ลิมิตที่ 128 MB | ไม่จำกัดปริมาณแต่คุณภาพจะลดลงเมื่อ > 1 MB


## Use Case Redis and RabbitMQ as Messaging Brokers
##### เปรียบเทียบ use case ของ Redis และ RabbitMQ เมื่อใช้โหมด messaging broker 
Feature | Redis | RabbitMQ
------------ | ------------- | -------------
Real-time chat applications | Best choice (low latency, simple pub/sub)(สมควรอย่างยิ่ง) | Overkill(เกินกำลัง)
Task queues / Job processing | เป็นไปไม่ได้ถ้าใช้ Pub/Sub ของ Redis | Best choice (message durability, retries)(สมการอย่างยิ่ง)
IoT messaging | Good for real-time updates(เหมาะกับงานที่ต้องการการอัพเดตแบบ real-time) | Good for persistent data and guaranteed delivery(เหมาะกับงานที่ต้องการความยั่งยืนของข้อมูล)
Transactional messaging | No built-in transaction support(ไม่เหมาะกับการทำงานแบบ Transaction) | Strong transactional guarantees(เหมาะมาก เพราะสามารถใช้กับการทำงานที่ complex logic)

## วิธีการเลือกใช้ระหว่าง Redis and RabbitMQ as Messaging Brokers
- ✅ Choose Redis if:
- ถ้าต้องการเลือกใช้ pub/sub ที่ส่งข้อมูลเร็ว เช่น live chat, real-time notifications
- ถ้าไม่ต้องการใช้กับงานที่ต้องตอบรับการส่งข้อความ(Ack) และความยั่งยืนของข้อมูล(persistence) 
- ถ้าต้องการใช้สำหรับงานที่ไม่ซับซ้อนของข้อมูล(lightweight messaging)
- ถ้าต้องการใช้กับงานที่อยาก set up ง่ายและ low latency.
- -------------------------------------------------
- ✅ Choose RabbitMQ if:
- ถ้าต้องการใช้กับข้อมูลที่เชื่อถือได้และยั่งยืน เช่น job queues, transaction logs
- ถ้าต้องการใช้กับงานที่ต้องตอบรับการส่งข้อความ(Ack) และ retry
- ถ้าต้องการใช้กับงานซับซ้อน โดยมีฟีเจอร์มาช่วย ได้แก่ fanout, topic-based filtering
- ถ้าต้องการใช้กับข้อมูลที่กลัวสูญหายหากเกิดข้อผิดพลาด
