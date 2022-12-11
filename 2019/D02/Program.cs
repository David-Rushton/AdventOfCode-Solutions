using System.IO;


const int Noun = 1;
const int Verb = 2;


// Star one
var one = RunOpComputer(12, 2);
Console.WriteLine($"Value at position 0: {one}");

// Star two
const int Target = 19690720;
for (var n = 0; n < 100; n++)
{
    for (var v = 0; v < 100; v++)
    {
        if (RunOpComputer(n, v) == Target)
        {
            Console.WriteLine($"Noun = {n} and verb = {v}: {100 * n + v}");
        }
    }
}


int RunOpComputer(int noun, int verb)
{
    var items = GetItems();
    items[Noun] = noun;
    items[Verb] = verb;

    for (var i = 0; i < items.Length; i += 4)
    {
        switch ((OpCode)items[i])
        {
            case OpCode.Add:
                items[items[i + 3]] = items[items[i + 1]] + items[items[i + 2]];
                break;

            case OpCode.Multiply:
                items[items[i + 3]] = items[items[i + 1]] * items[items[i + 2]];
                break;

            case OpCode.Halt:
                i = items.Length;
                continue;

            default:
                throw new Exception($"Unsupported op-code: {items[i]}");
        }
    }
    return items[0];
}


int[] GetItems()
{
    List<int> result = new();

    foreach (var line in File.ReadAllLines("./Input.txt"))
        foreach (var item in line.Split(','))
            if (int.TryParse(item, out var number))
                result.Add(number);

    return result.ToArray();
}

public enum OpCode
{
    Add = 1,
    Multiply = 2,
    Halt = 99
};
