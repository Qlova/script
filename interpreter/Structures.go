package interpreter

import "reflect"
import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Index(structure, index language.Type) language.Type {

	//Deterministic option.
	switch t := structure.(type) {
	case Array:

		var structure = implementation.RegisterOf(structure)
		var index = implementation.RegisterOf(index)
		switch t.Subtype.(type) {
		case Integer:
			var register = implementation.ReserveRegister()
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				var array = thread.Get(structure).([]int)
				thread.Set(register, array[thread.Get(index).(int)%len(array)])
			})
			return Integer{Expression: language.Statement(strconv.Itoa(register))}

			//TODO create reflect version.

		}

	case List:

		var structure = implementation.RegisterOf(structure)
		var index = implementation.RegisterOf(index)
		var register = implementation.ReserveRegister()

		implementation.AddInstruction(func(thread *dynamic.Thread) {
			thread.Set(register, reflect.ValueOf(thread.Get(structure)).Index(thread.Get(index).(int)).Interface())
		})
		return t.Subtype.Register(strconv.Itoa(register))
	}

	panic(implementation.Name() + ".Index(" + structure.Name() + ", " + index.Name() + ") Unimplemented")
	return nil
}

func (implementation Implementation) Modify(structure, index, value language.Type) language.Statement {

	switch t := structure.(type) {
	case Array:

		var structure = implementation.RegisterOf(structure)
		var index = implementation.RegisterOf(index)
		var value = implementation.RegisterOf(value)

		switch t.Subtype.(type) {
		case Integer:
			implementation.AddInstruction(func(thread *dynamic.Thread) {
				var array = thread.Get(structure).([]int)
				array[thread.Get(index).(int)%len(array)] = thread.Get(value).(int)
			})
			return language.Statement("")

			//TODO create reflect version.
		}

	case List:

		var structure = implementation.RegisterOf(structure)
		var index = implementation.RegisterOf(index)
		var value = implementation.RegisterOf(value)

		implementation.AddInstruction(func(thread *dynamic.Thread) {
			reflect.ValueOf(thread.Get(structure)).Index(thread.Get(index).(int)).Set(reflect.ValueOf(thread.Get(value)))
		})
		return language.Statement("")
	}

	panic(implementation.Name() + ".Modify() Unimplemented")
	return language.Statement("")
}
