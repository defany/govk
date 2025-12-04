# 😎 Очень крутой пакет для VK на Go

#### `наверное...`

### 💋 Основано на таких библиотеках, как:
VK sdk (Для golang): https://github.com/SevereCloud/vksdk <br/>
VK io (type script): https://github.com/negezor/vk-io <br/>
Telego (telegram + golang): https://github.com/mymmrac/telego

Мы стараемся взять лучшее из этих библиотек и воплотить в одной!

# 🤩 Возможности

## 😻 Удобная обработка команд от пользователей:

Создайте хендлер при помощи встроенного hear manager'a, используйте встроенные правила валидации запроса или напишите свои.
```go
	// making a new handler
	handler := hearManager.NewHandler(func(ctx context.Context, event msgmodel.MessagesNew) {
		params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а!")

		params.WithPeerID(event.Message.PeerID)

		_, err := vk.Api.Messages.Send(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// builtin match validators
	handler.WithMatchRules(heargo.WithMatchWord("123"), heargo.WithWordsIn("hello", "hi"))
```

## ⚒️ Удобство разработки
Забудьте о проблемах с типизацией в го и с неудобными параметрами. Пакет позволяет вам отправлять запрос так, как вы хотите
и указывать те параметры, которые вам нужны.

Необходимо отправить запрос с другой версией API? Пожалуйста! 

Хотите обрабатывать longpoll ивенты со своей типизацией? Просто укажите свою структуру в качестве дженерика

Библиотека старается дать вам максимально возможный выбор для отправки запросов

Пример запросов:
```go
// Самый легкий вариант
params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а!")

params.WithPeerID(event.Message.PeerID)

_, err := vk.Api.Messages.Send(params)
if err != nil {
	log.Fatalln(err)
}

// Другой вариант
_, err = vk.Api.Messages.Send(api.Params{
    PeerIDsParam: []int{event.Message.PeerID},
	MessageParam: "Приветик!",
})
if err != nil { 
	log.Fatalln(err)
}

// Еще один вариант! В данном случае вы можете указать то, что вы хотите получить в ответ
req := api.NewRequest[[]MessagesSendResponse](vk.Api.Api)

res, err := req.Execute("messages.send", params)
if err != nil {
	return nil, err
}

log.Print(res)
```

## 🔔 Обработка ивентов

Мы дали возможность легко обрабатывать ивенты избегая any для ивентов.

Просто вызовите функцию: <br/>
1) `updates.On` для обработки событий <br/>
2) `updates.Use` для добавления middleware для событий

Пример:
```go
hearManager := heargo.NewManager(vk)

updates.Use(vk.Updates, "message_new", hearManager.Middleware)

updates.On(vk.Updates, "message_new", func(_ context.Context, msg msgmodel.MessagesNew) {
	params := requests.NewSendParams().
		WithMessage("Привет! Я ответил на твое сообщение с ID: %d", msg.Message.ID)

	params.WithPeerID(222856843)

	_, err := vk.Api.Messages.Send(params)
	if err != nil {
		log.Fatalln(err)
	}
})
```

В библиотеке нету нужной структуры для ивента? Не беда, вы всегда можете прописать свою структуру! ~~Или мапу~~

Учтите, что поля у структуры должны быть экспортируемыми :)
```go
updates.On[YourOwnStr](vk.Updates, "message_new", func(_ context.Context, msg YourOwnStr) {
	log.Print("lets play!")
})
```

># 📜 Лицензия
### Проект распространяется под лицензией MIT. Подробнее см. LICENSE.

# 🆘 Контрибьюторство

Наша библиотека не имеет большого количества типов и методов, пожалуйста, если вам не хватает какого-то метода, 
создайте issue с запросом на создание или pull request с введением онного метода! 

~~Make VK+Go great again...~~
