package validator_test

import (

)

// type Person struct {
// 	Name  string `validate:"required"`
// 	Email string `validate:"required,email"`
// 	Age   int    `validate:"gte=18,lte=100"`
// }

// type Address struct {
// 	City   string `validate:"required"`
// 	Street string `validate:"required"`
// 	Zip    int    `validate:"required,numeric"`
// }

// type Product struct {
// 	Name     string  `validate:"required"`
// 	Price    float64 `validate:"gte=0"`
// 	Quantity int     `validate:"gte=0"`
// }

// func TestValidator_ValidStructs(t *testing.T) {
// 	v := validator.NewValidator()

// 	validPerson := Person{Name: "lewis Doe", Email: "lewis@example.com", Age: 25}
// 	err := v.Validate(validPerson)
// 	assert.Nil(t, err, "Expected no error for valid Person struct")

// 	validAddress := Address{City: "New York", Street: "123 Main St", Zip: 10001}
// 	err = v.Validate(validAddress)
// 	assert.Nil(t, err, "Expected no error for valid Address struct")

// 	validProduct := Product{Name: "Laptop", Price: 999.99, Quantity: 10}
// 	err = v.Validate(validProduct)
// 	assert.Nil(t, err, "Expected no error for valid Product struct")
// }

// func TestValidator_InvalidStructs(t *testing.T) {
// 	v := validator.NewValidator()

// 	invalidPerson := Person{Name: "", Email: "invalid-email", Age: 17}
// 	err := v.Validate(invalidPerson)
// 	assert.NotNil(t, err, "Expected validation error for invalid Person struct")

// 	invalidAddress := Address{City: "", Street: "123 Main St", Zip: 0}
// 	err = v.Validate(invalidAddress)
// 	assert.NotNil(t, err, "Expected validation error for invalid Address struct")

// 	invalidProduct := Product{Name: "", Price: -100.50, Quantity: -5}
// 	err = v.Validate(invalidProduct)
// 	assert.NotNil(t, err, "Expected validation error for invalid Product struct")
// }

// func TestValidator_EmptyStruct(t *testing.T) {
// 	v := validator.NewValidator()

// 	emptyPerson := Person{}
// 	err := v.Validate(emptyPerson)
// 	assert.NotNil(t, err, "Expected validation error for empty Person struct")

// 	emptyAddress := Address{}
// 	err = v.Validate(emptyAddress)
// 	assert.NotNil(t, err, "Expected validation error for empty Address struct")

// 	emptyProduct := Product{}
// 	err = v.Validate(emptyProduct)
// 	assert.NotNil(t, err, "Expected validation error for empty Product struct")
// }

// func TestValidator_CustomErrors(t *testing.T) {
// 	v := validator.NewValidator()

// 	outOfRangePerson := Person{Name: "lewis Doe", Email: "lewis@example.com", Age: 120}
// 	err := v.Validate(outOfRangePerson)
// 	assert.NotNil(t, err, "Expected validation error for out-of-range Age")

// 	invalidProduct := Product{Name: "Tablet", Price: -20, Quantity: -2}
// 	err = v.Validate(invalidProduct)
// 	assert.NotNil(t, err, "Expected validation error for negative Price and Quantity")
// }
