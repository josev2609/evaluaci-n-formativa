package main

import ("errors"
	"fmt"
	"time"
)


type Medicine struct {name string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int 
}


func NewMedicine(name, manufacturer string, manufactureDate time.Time, shelfLife int) (Medicine, error) {
	m := Medicine{}
	if err := m.SetName(name); err != nil {
		return Medicine{}, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return Medicine{}, err
	}
	if err := m.SetManufactureDate(manufactureDate); err != nil {
		return Medicine{}, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return Medicine{}, err
	}
	return m, nil
}


func (m *Medicine) Name() string {
	return m.name
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m *Medicine) Manufacturer() string {
	return m.manufacturer
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("el fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) ManufactureDate() time.Time {
	return m.manufactureDate
}

func (m *Medicine) SetManufactureDate(manufactureDate time.Time) error {
	if manufactureDate.IsZero() {
		return errors.New("la fecha de fabricación no es válida")
	}
	m.manufactureDate = manufactureDate
	return nil
}

func (m *Medicine) ShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("la vida útil debe ser un número positivo mayor a 0")
	}
	m.shelfLife = shelfLife
	return nil
}


func (m Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}


type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(name, manufacturer string, manufactureDate time.Time, shelfLife int, dosePerTablet float64, isPrescription bool) (Tablet, error) {
	baseMed, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return Tablet{}, err
	}
	t := Tablet{Medicine: baseMed, isPrescription: isPrescription}
	if err := t.SetDosePerTablet(dosePerTablet); err != nil {
		return Tablet{}, err
	}
	return t, nil
}

func (t *Tablet) DosePerTablet() float64 {
	return t.dosePerTablet
}

func (t *Tablet) SetDosePerTablet(dose float64) error {
	if dose <= 0 {
		return errors.New("la dosis por tableta debe ser mayor a 0")
	}
	t.dosePerTablet = dose
	return nil
}

func (t Tablet) DisplayDetails() {
	fmt.Printf("Tableta")
	fmt.Printf("Nombre:", t.Name())
	fmt.Printf("Fabricante:", t.Manufacturer())
	fmt.Printf("Fecha de Fabricación:", t.ManufactureDate().Format("2006-01-02"))
	fmt.Printf("Vida Útil:", t.ShelfLife())
	fmt.Printf("Dosis: mg", t.dosePerTablet)
	fmt.Printf("Requiere Receta: ", t.isPrescription)
}


type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewSyrup(name, manufacturer string, manufactureDate time.Time, shelfLife int, volume float64, flavor string) (Syrup, error) {
	baseMed, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return Syrup{}, err
	}
	s := Syrup{Medicine: baseMed, flavor: flavor}
	if err := s.SetVolume(volume); err != nil {
		return Syrup{}, err
	}
	return s, nil
}

func (s *Syrup) Volume() float64 {
	return s.volume
}

func (s *Syrup) SetVolume(volume float64) error {
	if volume <= 0 {
		return errors.New("el volumen debe ser mayor a 0")
	}
	s.volume = volume
	return nil
}


func main() {
	tabletInventory := []Tablet{}
	syrupInventory := []Syrup{}

	date1 := time.Date(2023, time.January, 15, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)

	t1, _ := NewTablet("Paracetamol", "PharmaECU", date1, 24, 500.0, false)
	t2, _ := NewTablet("Ibuprofeno", "Ecuhealth", date2, 18, 400.0, true)
	tabletInventory = append(tabletInventory, t1, t2)

	s1, _ := NewSyrup("Amoxicilina", "Ecumed", date1, 12, 120.0, "Fresa")
	syrupInventory = append(syrupInventory, s1)

	fmt.Println("INVENTARIO")
	for _, t := range tabletInventory {
		t.DisplayDetails()
		fmt.Printf("Caducidad:", t.ExpirationDate().Format("2006-01-02"))
	}
	for _, s := range syrupInventory {
		s.DisplayDetails()
		fmt.Printf("Caducidad:", s.ExpirationDate().Format("2006-01-02"))
	}
}