# RabbitMQ

    RabbitMQ เป็นระบบที่ช่วยในการส่งข้อความหรือข้อมูลระหว่างโปรแกรมหรือบริการต่างๆ ในรูปแบบที่เรียกว่า "message broker". คิดเหมือนเป็นไปรษณีย์ที่ทำหน้าที่ส่งจดหมายไปยังผู้รับที่แน่นอนตามที่อยู่ที่ระบุไว้ โดยในที่นี้จดหมายคือข้อมูลหรือคำสั่ง, และผู้รับคือโปรแกรมหรือบริการที่ต้องการข้อมูลนั้นๆ

    RabbitMQ ช่วยให้การสื่อสารระหว่างโปรแกรมต่างๆ เป็นไปอย่างมีระเบียบ โดยแยกการส่งข้อมูลออกจากการรับข้อมูล ซึ่งทำให้โปรแกรมต่างๆ สามารถทำงานได้อย่างอิสระและไม่ต้องรอหรือขึ้นอยู่กับโปรแกรมอื่นๆ นอกจากนี้ยังช่วยลดปัญหาเมื่อระบบใดระบบหนึ่งมีปัญหา, ข้อมูลจะไม่หายไปเพราะมันจะถูกเก็บไว้ในคิว (queue) รอจนกว่าจะสามารถส่งต่อไปยังผู้รับได้.

### ในระบบ RabbitMQ หรือระบบข้อความทั่วไป, "publish" และ "consumer" เป็นสองส่วนสำคัญในการจัดการข้อมูล:

    Publish (ผู้ส่ง) - หมายถึงการที่โปรแกรมหนึ่งๆ ส่งข้อมูลหรือข้อความไปยัง RabbitMQ. ผู้ส่งนี้ไม่จำเป็นต้องรู้ว่าใครจะเป็นผู้รับข้อความนี้, เพียงแค่มั่นใจว่าข้อความได้ถูกส่งไปยังระบบแล้ว.

    Consumer (ผู้รับ) - หมายถึงโปรแกรมที่รับข้อความจาก RabbitMQ. ผู้รับจะดึงข้อมูลหรือข้อความที่ต้องการจาก RabbitMQ เมื่อพร้อมจะประมวลผลข้อมูลนั้น. ผู้รับสามารถมีได้มากกว่าหนึ่งโปรแกรม, และสามารถรับข้อมูลที่เหมือนกันหรือต่างกันตามที่ได้ตั้งค่าไว้.

    โดยทั่วไป, ผู้ส่งและผู้รับทำงานอิสระต่อกัน และสามารถทำงานได้โดยไม่ต้องรอหรือขึ้นอยู่กับฝ่ายตรงข้าม, ช่วยให้ระบบสามารถทำงานได้มีประสิทธิภาพและความยืดหยุ่นสูง.

### ในระบบ RabbitMQ หรือระบบข้อความทั่วไปที่ใช้ message broker, "queue" และ "exchange" เป็นสองส่วนที่มีบทบาทสำคัญ:

    Queue (คิว) - คือสถานที่เก็บข้อความชั่วคราวในระบบ. ข้อความที่ถูกส่งโดยผู้ส่งจะเข้ามาอยู่ในคิวนี้จนกว่าจะมีผู้รับพร้อมรับข้อความนั้นๆ ไปประมวลผล. คิวสามารถรองรับหลายข้อความพร้อมกัน และทำให้การส่งข้อมูลระหว่างผู้ส่งและผู้รับมีความยืดหยุ่นมากขึ้น เพราะผู้ส่งสามารถส่งข้อมูลได้ต่อเนื่องโดยไม่ต้องรอผู้รับ.

    Exchange (แลกเปลี่ยน) - คือกลไกใน RabbitMQ ที่ทำหน้าที่รับข้อความจากผู้ส่งแล้วกำหนดว่าข้อความนั้นๆ ควรจะไปยังคิวไหน. แลกเปลี่ยนทำหน้าที่เป็นตัวกลางที่จัดการกฎหรือเงื่อนไขในการส่งต่อข้อความ โดยมีหลายประเภทเช่น direct, fanout, topic, ซึ่งแต่ละประเภทจะมีวิธีการส่งข้อความไปยังคิวที่แตกต่างกัน.

        - Direct Exchange  ส่งข้อความไปยังคิวที่มีการผูกกับ routing key ที่ตรงกัน. เหมาะสำหรับกรณีที่ต้องการให้ข้อความจากผู้ส่งไปยังคิวที่แน่นอนโดยมีการระบุ key อย่างชัดเจน.
        - Fanout Exchange ส่งข้อความไปยังทุกคิวที่ผูกกับมันโดยไม่คำนึงถึง routing key. นี้เหมาะสำหรับกรณีการกระจายข้อมูลหรือเหตุการณ์ไปยังผู้ฟังหลายๆ ตัวโดยไม่ต้องการควบคุมเป็นพิเศษ.
        - Topic Exchange ส่งข้อความไปยังคิวที่มีการผูกกับ routing key ที่ตรงกับ pattern ที่กำหนด. ช่วยให้สามารถกำหนดรูปแบบของ routing key ที่ซับซ้อนขึ้นได้, ทำให้มีความยืดหยุ่นในการระบุว่าข้อความจะไปถึงคิวไหนบ้างตาม pattern ที่ระบุ.
        - Headers Exchange ส่งข้อความไปยังคิวที่มีการผูกโดยอาศัย header ของข้อความมากกว่าเป็น routing key. ประเภทนี้ใช้ลักษณะของข้อมูลภายใน header ของข้อความเพื่อตัดสินใจว่าจะส่งข้อความไปยังคิวไหน.
        
        แต่ละประเภทของ Exchange ให้ความสามารถในการกระจายและควบคุมการส่งข้อความไปยังคิวต่างๆ ในวิธีที่ต่างกัน, เพื่อตอบสนองความต้องการที่หลากหลายของการใช้งานในระบบขนาดใหญ่.

    สรุปง่ายๆ คือ queue เหมือนกับกล่องจดหมายที่รอรับจดหมาย, และ exchange เหมือนเป็นบุคคลที่รับจดหมายจากผู้ส่งแล้วตัดสินใจว่าจะนำจดหมายนั้นไปใส่ในกล่องจดหมายไหนตามกฎที่กำหนดไว้.

## Code explain

### ในภาษาโปรแกรม Go ที่ใช้กับ RabbitMQ, ฟังก์ชันต่างๆ มีหน้าที่ดังนี้:

    exchangeDeclare() - ฟังก์ชันนี้ใช้สำหรับการสร้างหรือประกาศ "exchange" ใน RabbitMQ. Exchange คือตัวกลางที่จะรับข้อความจากผู้ส่งแล้วจัดการกระจายต่อไปยังคิวต่างๆ ตามกฎหรือเงื่อนไขที่กำหนดไว้ ซึ่งช่วยในการจัดการการส่งข้อความได้อย่างมีระเบียบและชัดเจน.

    queueDeclare() - ฟังก์ชันนี้ใช้สำหรับสร้างหรือประกาศ "queue" ใน RabbitMQ. Queue คือสถานที่เก็บข้อความชั่วคราวที่ข้อความจะถูกเก็บรอจนกว่าจะมีผู้รับพร้อมที่จะรับและประมวลผล. การประกาศคิวจะทำให้แน่ใจว่ามีที่เก็บสำหรับข้อความที่ส่งมาจากผู้ส่ง.

    queueBind() - ฟังก์ชันนี้ใช้เพื่อผูกหรือเชื่อมต่อ queue ที่ได้ประกาศไว้กับ exchange ที่สร้างขึ้น. การผูกนี้ช่วยให้ระบุได้ว่าข้อความที่ผ่านมาจาก exchange ควรจะไปยัง queue ใดโดยอาศัยเงื่อนไขหรือกฎการกระจายข้อความที่กำหนดไว้.
    เหล่านี้คือฟังก์ชันพื้นฐานในการจัดการกับข้อความในระบบ RabbitMQ โดยใช้ภาษา Go, แต่ละฟังก์ชันมีบทบาทสำคัญในการรับส่งและจัดการข้อความให้เป็นไปอย่างเป็นระเบียบและมีประสิทธิภาพ.

#### การเรียกใช้ฟังก์ชัน RabbitmqPublish:

    Rabbitmq.RabbitmqPublish(
        os.Getenv("EXCHANGE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_DOG"), // exchange name
        "", // routing key
        amqp.Publishing{
            Timestamp:   time.Now(),
            ContentType: "application/json",
            Body:        data,
        }
    )

`Exchange Name`: ชื่อของ exchange ที่จะส่งข้อความไปยังมัน. ชื่อ exchange นี้ได้มาจากตัวแปรสภาพแวดล้อม EXCHANGE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_DOG, ซึ่งเป็นการกำหนดไว้ว่าข้อความควรจะถูกส่งไปที่ exchange ไหน.
`Routing Key`: ในโค้ดนี้ routing key ถูกกำหนดเป็นสตริงว่าง, ซึ่งแสดงถึงการส่งข้อความไปยังทุกคิวที่ผูกกับ exchange นี้โดยไม่มีการกรองตาม key.
amqp.Publishing: โครงสร้างข้อมูลที่กำหนดวิธีการส่งข้อความ. รวมถึง:
`Timestamp`: ตั้งเวลาที่ข้อความถูกส่ง, ใช้ time.Now() ซึ่งเป็นเวลาปัจจุบัน.
ContentType: ระบุประเภทของข้อมูลที่ส่ง, ที่นี่เป็น "application/json", บ่งบอกว่าข้อมูลที่ส่งเป็นข้อมูลแบบ JSON.
`Body`: ข้อมูลที่จะส่ง. ในกรณีนี้, data คือตัวแปรที่มีข้อมูลที่จะถูกส่ง.

โดยรวมแล้ว, ฟังก์ชัน RabbitmqPublish นี้ใช้สำหรับการส่งข้อความไปยัง exchange โดยเฉพาะอย่างใน RabbitMQ, โดยข้อความนี้จะสามารถถูกส่งต่อไปยังคิวต่างๆ ตามการผูกที่ได้ตั้งค่าไว้กับ exchange นั้นๆ. การกำหนดให้ content type เป็น JSON ช่วยให้ผู้รับทราบถึงรูปแบบข้อมูลและสามารถจัดการข้อมูลนั้นได้ง่ายขึ้น.

#### การเรียกใช้งานฟังก์ชัน:

    go conn.RabbitmqHandleConsumedDeliveries(serviceName, "", delivery, messageHandler)

ฟังก์ชัน RabbitmqHandleConsumedDeliveries ถูกเรียกใช้งานโดยใช้ goroutine (คำว่า go ข้างหน้า) ซึ่งเป็นวิธีที่ Go ทำงานแบบ concurrent หรือพร้อมกัน. ฟังก์ชันนี้เริ่มต้นกระบวนการรับและจัดการข้อความจาก RabbitMQ.

#### ฟังก์ชัน RabbitmqHandleConsumedDeliveries:

    func (c *Connection) RabbitmqHandleConsumedDeliveries(queue, consumer string, delivery <-chan amqp.Delivery, fn func(*Connection, <-chan amqp.Delivery)) {
        for {
            go fn(c, delivery)

            if err := <-c.err; err != nil {
                for {
                    err := c.reconnect()
                    if err == nil {
                        break
                    }
                }

                delivery, err = c.RabbitmqConsume(queue, consumer)
                if err != nil {
                    log.Panic("failed to consume")
                }
            }
        }
    }

ฟังก์ชันนี้คอยตรวจจับและจัดการกับข้อความผ่าน channel delivery. ถ้าเกิดข้อผิดพลาดในการเชื่อมต่อ, มันจะพยายามเชื่อมต่อใหม่และเริ่มการรับข้อความอีกครั้ง. การทำงานนี้จะวนลูปตลอดเวลาด้วยการใช้ for.

#### ฟังก์ชัน messageHandler:

    func messageHandler(c *helpers.Connection, deliveries <-chan amqp.Delivery) {

        for d := range deliveries {
            if os.Getenv("SERVICE_NAME") == os.Getenv("QUEUE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_CAT") {
                worker.PublishMessageToConsumerCat(d.Body)
            } else if os.Getenv("SERVICE_NAME") == os.Getenv("QUEUE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_DOG") {
                worker.PublishMessageToConsumerDog(d.Body)
            }

            d.Ack(false)
        }
    }

ฟังก์ชันนี้รับข้อความจาก channel deliveries และตรวจสอบเงื่อนไขของแต่ละข้อความเพื่อทำงานที่เหมาะสม (เช่น ส่งข้อความไปยังบริการหรือฟังก์ชันที่เกี่ยวข้อง). ทุกข้อความที่ได้รับจะได้รับการยืนยันว่าได้รับแล้ว (d.Ack(false)).


#### ฟังก์ชัน RabbitmqConsume:

    func (c *Connection) RabbitmqConsume(queue, consumer string) (<-chan amqp.Delivery, error) {
        delivery, err := c.channel.Consume(
            queue,    //queue
            consumer, //consumer
            false,    //autoAck
            false,    //exclusive
            false,    //noLocal
            false,    //noWait
            nil,      //args
        )
        if err != nil {
            log.Printf("failed to consume: %s", err)
            return nil, err
        }

        return delivery, nil
    }

`queue`: ชื่อของ queue ที่จะรับข้อความจากมัน. ข้อความที่อยู่ใน queue นี้จะถูกส่งไปยัง consumer ที่เชื่อมต่อกับ queue นี้.  
`consumer`: ตัวระบุของ consumer ที่ใช้รับข้อความ. สามารถใช้เพื่อระบุ consumer เฉพาะหรือใช้เป็นที่เก็บสถิติ.  
`autoAck` (Automatic Acknowledgement): ถ้าตั้งค่าเป็น true, ข้อความที่ได้รับจะถูกยืนยันอัตโนมัติว่าได้รับแล้ว (acknowledged) ซึ่งหมายความว่าระบบจะถือว่าข้อความนั้นได้รับการจัดการเรียบร้อยแล้วเมื่อมันถูกส่งไปยัง consumer. ถ้าตั้งเป็น false,   จะต้องมีการยืนยันข้อความด้วยตัวเองเพื่อบอกว่าได้รับและจัดการข้อความนั้นเรียบร้อยแล้ว.  
`exclusive`: ถ้าตั้งค่าเป็น true, queue นี้จะเป็นของ consumer นี้เท่านั้นที่สามารถเข้าถึงได้. นี่ช่วยป้องกันไม่ให้ consumer อื่นเข้าถึงข้อความใน queue นี้.  
`noLocal`: ถ้าตั้งค่าเป็น true, ข้อความที่ส่งโดย connection นี้จะไม่ถูกส่งกลับมายัง consumer นี้. โดยปกติ, ตัวเลือกนี้จะใช้ไม่บ่อยในส่วนใหญ่ของการตั้งค่า   RabbitMQ.  
`noWait`: ถ้าตั้งค่าเป็น true, การดำเนินการจะไม่รอการตอบกลับจากเซิร์ฟเวอร์ การตั้งค่านี้สามารถช่วยลดเวลาในการตอบสนองได้ในสถานการณ์บางอย่าง  แต่จะสูญเสียการรับประกันว่าเซิร์ฟเวอร์ได้ตอบรับการตั้งค่าได้สำเร็จแล้ว.  
`args`: อาร์กิวเมนต์เพิ่มเติมในรูปแบบ key-value pairs ที่สามารถส่งผ่านไปยังเซิร์ฟเวอร์เพื่อกำหนดตัวเลือกการดำเนินการเพิ่มเติม.  

ฟังก์ชัน Consume นี้สำคัญมากในการสร้างการเชื่อมต่อระหว่างโปรแกรมของคุณกับ RabbitMQ เพื่อรับข้อความจาก queue อย่างต่อเนื่อง. การตั้งค่าพารามิเตอร์ต่างๆ  เหล่านี้จะมีผลต่อวิธีการรับและจัดการข้อความของคุณ.


## example using with request api

### consumer > exchange > publish-message-to-consumer-cat-exchange (ทำงานอยู่ที่ container name publish-message-to-consumer-cat)

    curl --location 'http://127.0.0.1:9000/api/v1/publish_message_to_consumer_cat' \
    --header 'Content-Type: application/json' \
    --data '{
        "message": "The cat doesn'\''t like eat milk.",
        "name": "tony",
        "age": 2
    }'

    หลังจากเรียก API ที่ /api/v1/publish_message_to_consumer_cat,
    ระบบจะทำการ publish ข้อมูลไปยัง exchange ที่ชื่อ publish-message-to-consumer-cat-exchange
    
    จากนั้น consumer จะดึงข้อมูลด้วย queue ที่ได้ผูกไว้กับ exchange นี้
    มาแสดง log ดังนี

    [CAT] name: tony, age: 2, message: The cat doesn't like eat milk.


### consumer > exchange > publish-message-to-consumer-dog-exchange (ทำงานอยู่ที่ container name publish-message-to-consumer-dog)

    curl --location 'http://127.0.0.1:9000/api/v1/publish_message_to_consumer_dog' \
    --header 'Content-Type: application/json' \
    --data '{
        "bleed": "chivava",
        "message": "chivava เป็นสายพันธุ์ของสุนัขที่มีขนาดเล็กที่สุดในโลก"
    }'

    หลังจากเรียก API ที่ /api/v1/publish_message_to_consumer_dog,
    ระบบจะทำการ publish ข้อมูลไปยัง exchange ที่ชื่อ publish-message-to-consumer-dog-exchange
    
    จากนั้น consumer จะดึงข้อมูลด้วย queue ที่ได้ผูกไว้กับ exchange นี้
    มาแสดง log ดังนี

    [DOG] breed: chivava, message: chivava เป็นสายพันธุ์ของสุนัขที่มีขนาดเล็กที่สุดในโลก