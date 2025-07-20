package main

import (
	"fmt"
	"time"
)

func main() {

	july := map[int]string{
		1:  "Bl. Junipero Serra",
		2:  "Sts. Processus and Martinian, Martyrs",
		3:  "Pope St. Leo & St. Irenaeus",
		4:  "St. Theodore",
		5:  "St. Anthony May Zaccaria",
		6:  "St. John Fisher & St. Thomas More",
		7:  "Sts. Cyril and Methodius",
		8:  "St. Elizabeth of Portugal",
		9:  "St. Maria Goretti",
		10: "Holy Seven Brothers with Sts. Rufina and Secunda, Martyrs",
		11: "Pope St. Pius I, Martyr | St. Oliver Plunket",
		12: "St. John Gualbert | Sts. Nabor and Felix, Martyrs | St. Veronica of the Veil",
		13: "Pope St. Anacletus, Martyr | St. Mildred | St. Teresa of the Andes",
		14: "St. Bonaventure | St. Francis Solano | Bl. Kateri Tekakwitha",
		15: "St. Henry II, Emperor",
		16: "No Official Saint Today!",
		17: "St. Alexis the Beggar",
		18: "St. Symphorosa and Her Children, Martyrs & St. Camillus de Lellis",
		19: "St. Vincent de Paul",
		20: "St. Jerome Emilian",
		21: "St. Laurence of Brindisi & St. Praxedes",
		22: "St. Mary Magdalen",
		23: "St. Apollinaris of Ravenna & St. Liborius, Martyrs",
		24: "St. Christina",
		25: "Vigil of St. James the Greater, Apostle | St. James the Greater, Apostle | St. Christopher",
		26: "St. Anne",
		27: "St. Pantaleon, Martyr | Pope St. Celestine I",
		28: "Sts. Nazarius and Celsus, Victor I, and Innocent I, Martyrs",
		29: "St. Martha & 20: Sts. Felix II, Simplicius, Faustinus, Beatrice",
		30: "Sts. Abdon and Sennen | St. Ignatius of Loyola",
	}

	currentTime := time.Now()
	current_month := currentTime.Month()
	current_date := currentTime.Day()

	// Catppuccin Colors
	// var Red = "\x1b[38;2;243;139;168m"
	var text = "\x1b[38;2;205;214;244m"

	// Main Logic
	switch current_month.String() {
	case "July":
		fmt.Println(text + july[current_date])
	}

}
