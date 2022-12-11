var test = args.Contains("--test");
var starTwo = args.Contains("--star-two");
var path = test ? "./input.test.txt" : "./input.txt";
var lines = File.ReadAllLines(path);

var polymerTemplate = lines[0] ?? throw new NullReferenceException();
var pairMap = new Dictionary<string, string>();
var pairCount = new Dictionary<string, long>();


foreach (var rule in lines.Skip(2))
{
    var elements = rule.Split(" -> ");
    pairMap.Add(elements[0], $"{elements[0][0]}{elements[1]}");
    pairCount.Add(elements[0], 0);
}


for (var i = 0; i < polymerTemplate.Length - 1; i++)
{
    pairCount[polymerTemplate[i..(i + 2)]]++;
}

// Console.WriteLine($"\nBefore");
// PrintProgress();


var steps = starTwo ? 41 : 11;
for (var step = 1; step < steps; step++)
{
    var next = new Dictionary<string, long>(pairCount);

    foreach (var pair in pairCount)
    {
        if (pair.Value > 0)
        {
            var pairX = pairMap[pair.Key];
            var pairY = $"{pairX[1]}{pair.Key[1]}";
            next[pair.Key] -= pair.Value;
            next[pairX] += pair.Value;
            next[pairY] += pair.Value;
        }
    }

    pairCount = next;

    // Console.WriteLine($"\nAfter step {step}");
    // PrintProgress();
}

PrintResult();


// void PrintProgress()
// {
//     _ = pairCount ?? throw new NullReferenceException(nameof(pairCount));

//     foreach (var pair in pairCount.Where(p => p.Value > 0).OrderBy(p => p.Key[1]))
//     {
//         Console.WriteLine($"{pair.Key} -> {pair.Key[1]}: {pair.Value} ");
//         Console.WriteLine(pairCount.Sum(p => p.Value));
//     }
// }

void PrintResult()
{
    var byLetter =
    (
        from pair in pairCount
        group pair by pair.Key[1] into g
        orderby g.Key
        select new
        {
            Character = g.Key,
            Value = g.Sum(kvp => kvp.Value)
        }
    ).ToDictionary(k => k.Character, v => v.Value);

    byLetter[polymerTemplate[0]]++;
    var min = byLetter.OrderBy(i => i.Value).First().Value;
    var max = byLetter.OrderByDescending(i => i.Value).First().Value;

    Console.WriteLine("Result-----------------");

    foreach (var result in byLetter.Where(l => l.Value > 0))
        Console.WriteLine($"{result.Key}: {result.Value}");

    Console.WriteLine($"\nScore: {max - min}\n");
}



































// Console.WriteLine($"Template:     {polymerTemplate}");

// var steps = starTwo ? 41 : 11;
// for (var step = 1; step < steps; step ++)
// {
//     var polymer = new List<char> { polymerTemplate[0] };

//     for (var i = 0; i < polymerTemplate.Count - 1; i++)
//     {
//         var pair = string.Join("", polymerTemplate.Skip(i).Take(2));
//         var insertion = pairInsertionRules[pair];

//         polymer.Add(insertion);
//         polymer.Add(pair[1]);
//     }

//     Console.WriteLine($"After step {step}: {polymer.Count}");
//     polymerTemplate = polymer;
// }


// var population = new Dictionary<char, long>();
// foreach (var c in polymerTemplate)
// {
//     if (population.ContainsKey(c))
//     {
//         population[c]++;
//     }
//     else
//     {
//         population.Add(c, 1);
//     }
// }

// var mostPopular = population.OrderByDescending(i => i.Value).First();
// var leastPopular = population.OrderBy(i => i.Value).First();

// Console.WriteLine($"Most common = {mostPopular.Key} ({mostPopular.Value}) | Least common = {leastPopular.Key} ({leastPopular.Value}) | Score = {mostPopular.Value - leastPopular.Value}");
