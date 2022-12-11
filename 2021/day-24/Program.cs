var path = args[0];
Queue<long> inputs = new Queue<long>();
long w = 0;
long x = 0;
long y = 0;
long z = 0;


for (long i = 99; i >= 1; i--)
{
    if (i.ToString().Contains('0'))
        continue;

    inputs = new Queue<long>(i.ToString().Select(i => long.Parse(i.ToString())));
    var result = RunProgram(path);

    if (result == 0)
    {
        Console.WriteLine("We made it! {i}");
        Environment.Exit(0);
    }
    else
    {
        Console.WriteLine($"{i} = {result} (w = {w}, x = {x}, y = {y}, z = {z})");
        //i -= result;
    }

    // if (i % 1000 == 1)
    // {
    //     Environment.Exit(0);
    // }
}


long RunProgram(string path)
{
    w = 0;
    x = 0;
    y = 0;
    z = 0;
    foreach (var instruction in File.ReadAllLines(path))
    {
        if (instruction == string.Empty)
            return -1;

        var elements = instruction.Split(" ");
        var operation = elements[0];
        var varA = elements[1];

        if (operation == "inp")
        {
            WriteVariable(varA, PromptVariable(varA));
        }
        else
        {
            var varB = elements[2];
            var a = ReadVariable(varA);
            var b = ReadVariable(varB);

            switch (operation)
            {
                case "add":
                    WriteVariable(varA, a + b);
                    break;

                case "mul":
                    WriteVariable(varA, a * b);
                    break;

                case "div":
                    WriteVariable(varA, a / b);
                    break;

                case "mod":
                    WriteVariable(varA, a % b);
                    break;

                case "eql":
                    WriteVariable(varA, a == b ? 1 : 0);
                    break;

                default:
                    throw new Exception($"Operation not supported {operation}");
            }
        }
    }

    // Console.WriteLine($"Result = {{ w = {w}, x = {x}, y = {y}, z = {z} }}");
    return z;
}



long PromptVariable(string variableName)
{
    if (inputs.Count > 0)
        return inputs.Dequeue();

    Console.Write($"Enter a integer for {variableName}: ");
    var value = Console.ReadLine();

    if (long.TryParse(value, out var number))
    {
        return number;
    }

    throw new Exception($"{value} is not a number");
}

long ReadVariable(string value)
{
    if (long.TryParse(value, out var number))
    {
        return number;
    }

    return (value) switch
    {
        "w" => w,
        "x" => x,
        "y" => y,
        "z" => z,
        _ => throw new Exception($"Variable not supported {value}")
    };
}

void WriteVariable(string variableName, long value)
{
    switch (variableName)
    {
        case "w":
            w = value;
            break;

        case "x":
            x = value;
            break;

        case "y":
            y = value;
            break;

        case "z":
            z = value;
            break;

        default:
            throw new Exception($"Variable not supported {variableName}");
    }
}
