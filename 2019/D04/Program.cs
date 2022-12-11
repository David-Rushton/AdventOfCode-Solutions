const int Start = 284639;
const int End = 748759;

var candidates = 0;
var candidatesTwo = 0;
for (var current = Start; current < End; current++)
{
    if (IsCandidate(current))
        candidates++;

    if (IsCandidateTwo(current))
        candidatesTwo++;
}


Console.WriteLine($"Star one: {candidates}");
Console.WriteLine($"Star one: {candidatesTwo}");

bool IsCandidate(int value)
{
    var valueStr = value.ToString().ToCharArray();
    bool hasAdjacent = false;

    for (var i = 0; i < valueStr.Length - 1; i++)
    {
        var left = char.GetNumericValue(valueStr[i]);
        var right = char.GetNumericValue(valueStr[i + 1]);


        if (left > right)
            return false;

        if (left == right)
            hasAdjacent = true;
    }

    return hasAdjacent;
}

bool IsCandidateTwo(int value)
{
    var valueStr = value.ToString();
    var adjacentCount = 0;
    bool pairFound = false;

    for (var i = 0; i < valueStr.Length - 1; i++)
    {
        var left = char.GetNumericValue(valueStr[i]);
        var right = char.GetNumericValue(valueStr[i + 1]);


        if (left > right)
            return false;

        if (left == right)
        {
            adjacentCount++;
        }
        else
        {
            if (adjacentCount == 2)
                pairFound = true;

            adjacentCount = 0;
        }
    }


    return pairFound;
}
