package dsl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var _ = Describe("Attribute", func() {
	var parent *AttributeDefinition
	var name string
	var dataType DataType
	var description string
	var dsl func()

	BeforeEach(func() {
		parent = new(AttributeDefinition)
		name = ""
		dataType = nil
		description = ""
		dsl = nil
	})

	JustBeforeEach(func() {
		Reset([]DSLDefinition{parent})
		if dsl == nil {
			if dataType == nil {
				Attribute(name)
			} else if description == "" {
				Attribute(name, dataType)
			} else {
				Attribute(name, dataType, description)
			}
		} else if dataType == nil {
			Attribute(name, dsl)
		} else if description == "" {
			Attribute(name, dataType, dsl)
		} else {
			Attribute(name, dataType, description, dsl)
		}
	})

	Context("with only a name", func() {
		BeforeEach(func() {
			name = "foo"
		})

		It("produces an attribute of type string", func() {
			t := parent.Type
			Ω(t).ShouldNot(BeNil())
			Ω(t).Should(BeAssignableToTypeOf(Object{}))
			o := t.(Object)
			Ω(o).Should(HaveLen(1))
			Ω(o).Should(HaveKey(name))
			Ω(o[name].Type).Should(Equal(String))
		})
	})

	Context("with a name and datatype", func() {
		BeforeEach(func() {
			name = "foo"
			dataType = Integer
		})

		It("produces an attribute of given type", func() {
			t := parent.Type
			Ω(t).ShouldNot(BeNil())
			Ω(t).Should(BeAssignableToTypeOf(Object{}))
			o := t.(Object)
			Ω(o).Should(HaveLen(1))
			Ω(o).Should(HaveKey(name))
			Ω(o[name].Type).Should(Equal(Integer))
		})
	})

	Context("with a name, datatype and description", func() {
		BeforeEach(func() {
			name = "foo"
			dataType = Integer
			description = "bar"
		})

		It("produces an attribute of given type and given description", func() {
			t := parent.Type
			Ω(t).ShouldNot(BeNil())
			Ω(t).Should(BeAssignableToTypeOf(Object{}))
			o := t.(Object)
			Ω(o).Should(HaveLen(1))
			Ω(o).Should(HaveKey(name))
			Ω(o[name].Type).Should(Equal(Integer))
			Ω(o[name].Description).Should(Equal(description))
		})
	})

	Context("with a name and a DSL defining an enum validation", func() {
		BeforeEach(func() {
			name = "foo"
			dsl = func() { Enum("one", "two") }
		})

		It("produces an attribute of type string with a validation", func() {
			t := parent.Type
			Ω(t).ShouldNot(BeNil())
			Ω(t).Should(BeAssignableToTypeOf(Object{}))
			o := t.(Object)
			Ω(o).Should(HaveLen(1))
			Ω(o).Should(HaveKey(name))
			Ω(o[name].Type).Should(Equal(String))
			Ω(o[name].Validations).Should(HaveLen(1))
			Ω(o[name].Validations[0]).Should(BeAssignableToTypeOf(&EnumValidationDefinition{}))
		})
	})

	Context("with a name, type integer and a DSL defining an enum validation", func() {
		BeforeEach(func() {
			name = "foo"
			dataType = Integer
			dsl = func() { Enum("one", "two") }
		})

		It("produces an attribute of type integer with a validation", func() {
			t := parent.Type
			Ω(t).ShouldNot(BeNil())
			Ω(t).Should(BeAssignableToTypeOf(Object{}))
			o := t.(Object)
			Ω(o).Should(HaveLen(1))
			Ω(o).Should(HaveKey(name))
			Ω(o[name].Type).Should(Equal(Integer))
			Ω(o[name].Validations).Should(HaveLen(1))
			Ω(o[name].Validations[0]).Should(BeAssignableToTypeOf(&EnumValidationDefinition{}))
		})
	})
	Context("with a name, type integer, a description and a DSL defining an enum validation", func() {
		BeforeEach(func() {
			name = "foo"
			dataType = Integer
			description = "bar"
			dsl = func() { Enum("one", "two") }
		})

		It("produces an attribute of type integer with a validation and the description", func() {
			t := parent.Type
			Ω(t).ShouldNot(BeNil())
			Ω(t).Should(BeAssignableToTypeOf(Object{}))
			o := t.(Object)
			Ω(o).Should(HaveLen(1))
			Ω(o).Should(HaveKey(name))
			Ω(o[name].Type).Should(Equal(Integer))
			Ω(o[name].Validations).Should(HaveLen(1))
			Ω(o[name].Validations[0]).Should(BeAssignableToTypeOf(&EnumValidationDefinition{}))
			Ω(o[name].Description).Should(Equal(description))
		})
	})
})