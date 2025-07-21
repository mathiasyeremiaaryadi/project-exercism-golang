package main

import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && (len(r.Address) != 0 && r.Address["street"] != "")
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = map[string]string(nil)
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	total := 0

	for _, v := range residents {
		if v.HasRequiredInfo() {
			total += 1
		}
	}

	return total
}

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	fmt.Println(NewResident(name, age, address))
	// => &{Matthew Sanabria 29 map[street:Main St.]}

	name = "Matthew Sanabria"
	age = 0
	address = make(map[string]string)

	resident := NewResident(name, age, address)

	resident.HasRequiredInfo()
	// => false
	fmt.Println(resident)

	name = "Matthew Sanabria"
	age = 29
	address = map[string]string{"street": "Main St."}

	resident = NewResident(name, age, address)

	fmt.Println(resident)
	// Output: &{Matthew Sanabria 29 map[street:Main St.]}

	resident.Delete()

	fmt.Println(resident)
	// Output: &{ 0 map[]}

	name1 := "Matthew Sanabria"
	age1 := 29
	address1 := map[string]string{"street": "Main St."}

	resident1 := NewResident(name1, age1, address1)

	name2 := "Rob Pike"
	age2 := 0
	address2 := make(map[string]string)

	resident2 := NewResident(name2, age2, address2)

	residents := []*Resident{resident1, resident2}

	fmt.Println(Count(residents))
	// => 1
}
