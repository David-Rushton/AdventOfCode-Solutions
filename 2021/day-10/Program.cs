using System.Text;

var test = args.Contains("--test");
var path = test ? "./input.test.txt" : "./input.txt";
var lines = File.ReadAllLines(path);
var illegalSyntaxScore = 0;
var autocompleteLines = new List<(int lineNumber, string correction)>();
var autocompleteScores = new List<long>();

for (var line = 0; line < lines.Length; line++)
{
    var stack = new Stack<char>();
    for (var character = 0; character < lines[line].Length; character++)
    {
        var current = lines[line][character];

        if (IsOpening(current))
        {
            stack.Push(current);
        }
        else
        {
            var opening = stack.Pop();
            var closing = GetClosingPartner(opening);

            if (current != closing)
            {
                Console.WriteLine($"Syntax error.  Expected {closing} but found {current} on line: {line + 1}, position {character + 1}.");
                illegalSyntaxScore += GetIllegalSyntaxScore(current);
                goto EndLine;
            }
        }
    }

    autocompleteLines.Add((line, GetAutocompleteCorrection(stack)));
// Console.WriteLine($"Autocompleting line {line + 1} with {autocomplete.correction} - {autocomplete.score} total points.");
// autocompleteScores.Add(autocomplete.score);

EndLine:
    _ = 1;
}


foreach (var autocomplete in autocompleteLines)
{
    var score = GetAutocompleteScore(autocomplete.correction, autocompleteLines.Count);
    autocompleteScores.Add(score);
    Console.WriteLine($"Autocompleting line {autocomplete.lineNumber} with {autocomplete.correction} - {score}");
}
var autocompleteScore = autocompleteScores
    .OrderBy(s => s)
    .Skip(autocompleteScores.Count / 2)
    .Take(1)
    .Single();


Console.WriteLine($"Star one | Illegal syntax score {illegalSyntaxScore}");
Console.WriteLine($"Star two | Autocomplete score {autocompleteScore}");


bool IsOpening(char toCheck) => toCheck is '(' or '[' or '{' or '<';

char GetClosingPartner(char openingCharacter) =>
    (openingCharacter) switch
    {
        '(' => ')',
        '[' => ']',
        '{' => '}',
        '<' => '>',
        _ => throw new Exception($"Unexpected characer {openingCharacter}")
    };

int GetIllegalSyntaxScore(char closingCharacter) =>
    (closingCharacter) switch
    {
        ')' => 3,
        ']' => 57,
        '}' => 1197,
        '>' => 25137,
        _ => throw new Exception($"Unexpected characer {closingCharacter}")
    };

string GetAutocompleteCorrection(Stack<char> stack)
{
    var sb = new StringBuilder();

    while (stack.Any())
        sb.Append(GetClosingPartner(stack.Pop()));

    return sb.ToString();
}


long GetAutocompleteScore(string correction, int multiplier)
{
    long score = 0;

    foreach (var character in correction)
    {
        score *= 5;
        score += (character) switch
        {
            ')' => 1,
            ']' => 2,
            '}' => 3,
            '>' => 4,
            _ => throw new Exception("Unexpected characer on stack")
        };
    }

    return score;
}
