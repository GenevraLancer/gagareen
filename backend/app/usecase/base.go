package entity

//Роль разработчика формуляров в Студии
type FormularDeveloper struct {
	checkSelection() //Подбор метода проверки из es
}

//Роль продакт менеджера Студии 
type ProductManager struct {
	viewSearchStats() //Посмотреть частоту использования разработчиками формуляров поискового движка
}

type Formular struct {
	check   //проверки
	project //проект, в котором разрабатыватся формуляр
}

type Check struct {
	methodID //имя бина и метода
	desc     //описание проверки
}
