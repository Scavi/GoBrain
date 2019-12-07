package aoc

type Mode int
type Operation int

const (
	PositionMode  Mode = 0
	ImmediateMode Mode = 1
)

const (
	Op1Add         Operation = 1
	Op2Multiply    Operation = 2
	Op3SaveInput   Operation = 3
	Op4SaveOutput  Operation = 4
	Op5JumpIfTrue  Operation = 5
	Op6JumpIfFalse Operation = 6
	Op7LessThan    Operation = 7
	Op8Equals      Operation = 8
	OpEnd          Operation = 99
)

type Instruction struct {
	OpCode     Operation
	ModeParam1 Mode
	ModeParam2 Mode
	ModeOutput Mode
}

func _instructionDetailsFrom(input int) Instruction {
	return Instruction{
		Operation(input % 100),
		Mode((input / 100) % 10),
		Mode((input / 1000) % 10),
		Mode(input / 10000),
	}
}

func _valueOf(input []int, paramMode Mode, lookupPos int) int {
	value := 0
	switch paramMode {
	case PositionMode:
		value = input[input[lookupPos]]
	case ImmediateMode:
		value = input[lookupPos]
	default:
		panic("What an evil mode!")
	}
	return value
}

func _set(input [] int, pos int, value int) {
	input[pos] = value
	print(value)
	print("\t")
	print(pos)
	print("\t")
}

func _add(input [] int, pos int, instruction Instruction) (int, int) {
	result := _valueOf(input, instruction.ModeParam1, pos+1) +
		_valueOf(input, instruction.ModeParam2, pos+2)
	print("add\t\t")
	print("\t")
	_set(input, input[pos+3], result)
	print(pos + 4)
	print("\n")
	return result, pos + 4
}

func _multiply(input [] int, pos int, instruction Instruction) (int, int) {
	result := _valueOf(input, instruction.ModeParam1, pos+1) *
		_valueOf(input, instruction.ModeParam2, pos+2)
	print("multiply\t\t")
	_set(input, input[pos+3], result)
	print(pos + 4)
	print("\n")
	return result, pos + 4
}

func _input(input []int, pos int, inputValue int) int {
	print("_input\t\t")
	_set(input, input[pos+1], inputValue)
	print(pos + 2)
	print("\n")
	return pos + 2
}

func _output(input []int, pos int, instruction Instruction) (int, int) {
	// 72, 104, 126, 170, 204, 333, 363, 498, 513, 543, 603, 618, 633
	result := _valueOf(input, instruction.ModeParam1, pos+1)
	// result := input[pos + 1]
	print("_print\t\t")

	print(result)
	print("\t")
	print(pos + 2)

	print("\n")
	print("Validate -> ")
	print(result)
	print("\n")
	return result, pos + 2
}

func _jumpIfTrue(input []int, pos int, instruction Instruction) int {
	param1 := _valueOf(input, instruction.ModeParam1, pos+1)
	newPos := pos
	if param1 != 0 {
		newPos = _valueOf(input, instruction.ModeParam2, pos+2)
	} else {
		newPos += 3
	}
	print("jit\t\t\t0\t")
	print(pos)
	print("\n")
	return newPos
}

func _jumpIfFalse(input []int, pos int, instruction Instruction) int {
	param1 := _valueOf(input, instruction.ModeParam1, pos+1)
	if param1 == 0 {
		pos = _valueOf(input, instruction.ModeParam2, pos+2)
	} else {
		pos += 3
	}
	return pos
}

func _lessThan(input []int, pos int, instruction Instruction) int {
	param1 := _valueOf(input, instruction.ModeParam1, pos+1)
	param2 := _valueOf(input, instruction.ModeParam2, pos+2)
	if param1 < param2 {
		_set(input, input[pos+3], 1)
	} else {
		_set(input, input[pos+3], 0)
	}
	return pos + 4
}

func _equals(input []int, pos int, instruction Instruction) int {
	param1 := _valueOf(input, instruction.ModeParam1, pos+1)
	param2 := _valueOf(input, instruction.ModeParam2, pos+2)
	if param1 == param2 {
		_set(input, input[pos+3], 1)
	} else {
		_set(input, input[pos+3], 0)
	}
	return pos + 4
}

func executeInstructions(input []int, inputValue int) int {
	run := true
	pos := 0
	result := 0
	for run {
		t := _instructionDetailsFrom(1005)
		instruction := _instructionDetailsFrom(input[pos])
		if t == instruction {
			print("")
		}
		switch instruction.OpCode {
		case Op1Add:
			result, pos = _add(input, pos, instruction)
		case Op2Multiply:
			result, pos = _multiply(input, pos, instruction)
		case Op3SaveInput:
			pos = _input(input, pos, inputValue)
		case Op4SaveOutput:
			result, pos = _output(input, pos, instruction)
		case Op5JumpIfTrue:
			pos = _jumpIfTrue(input, pos, instruction)
		case Op6JumpIfFalse:
			pos = _jumpIfFalse(input, pos, instruction)
		case Op7LessThan:
			pos = _lessThan(input, pos, instruction)
		case Op8Equals:
			pos = _equals(input, pos, instruction)
		case OpEnd:
			run = false
		default:
			panic("Oh no! What a weired op code!")
		}
	}
	return result
}
