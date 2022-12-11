/*
    0:      1:      2:      3:      4:       5:      6:      7:      8:      9:
    aaaa    ....    aaaa    aaaa    ....    aaaa    aaaa    aaaa    aaaa    aaaa
    b    c  .    c  .    c  .    c  b    c  b    .  b    .  .    c  b    c  b    c
    b    c  .    c  .    c  .    c  b    c  b    .  b    .  .    c  b    c  b    c
    ....    ....    dddd    dddd    dddd    dddd    dddd    ....    dddd    dddd
    e    f  .    f  e    .  .    f  .    f  .    f  e    f  .    f  e    f  .    f
    e    f  .    f  e    .  .    f  .    f  .    f  e    f  .    f  e    f  .    f
    gggg    ....    gggg    gggg    ....    gggg    gggg    ....    gggg    gggg

    abcefg  cf      acdeg   acdfg   bcdf    abdfg   abdefg  acf     abcdefg abcdfg
    6       2       5       5       4       5       6       3       7       6


    2 1     1
    3 1     7
    4 1     4
    5 3     2, 3 & 5
    6 3     6, 0 and 9
    7 1     8

    1 not in 7 == a
    4 not in 1 == (bd)

    1 =    c  f   2 *
    7 =  a c  f   3 *
    4 =   bcd f   4 *
    2 =  a cde g  5
    3 =  a cd fg  5
    5 =  abcd fg  6
    6 =  ab defg  6
    0 =  abc efg  6
    9 =  abcd fg  6
    8 =  abcdefg  7 *
*/


var starTwo = args.Contains("--star-two");
var test = args.Contains("--test");
var path = test ? "./Input.Test.txt" : "./Input.txt";
var lines = File.ReadAllLines(path);
var starOneAppearanceRate = 0;
var segments = Day08.SevenSegmentDigit.Build();
long starTwoValue = 0;


foreach (var line in lines)
{
    var elements = line.Split(" | ");
    var samples = elements[0].Split(' ');
    var outputs = elements[1].Split(' ');
    var appearanceRates = new Dictionary<char, int>
    {
        {'a', 0}, {'b', 0}, {'c', 0}, {'d', 0}, {'e', 0}, {'f', 0}, {'g', 0}
    };

    if (outputs.Length != 4)
        throw new Exception("Unexpected output length");

    if (samples.Length != 10)
        throw new Exception("Unexpected sample length");


    foreach (var output in outputs)
    {
        if (new[] { 2, 4, 3, 7 }.Contains(output.Length))
        {
            starOneAppearanceRate++;
        }
    }

    foreach (var sample in samples)
    {
        foreach (var character in sample)
        {
            appearanceRates[character]++;
        }
    }


    var segmentMap = new Dictionary<char, char>();
    var map = new Dictionary<int, string>();

    // Maps digits with unique lengths.
    map.Add(1, samples.Where(o => o.Length == 2).First());
    map.Add(4, samples.Where(o => o.Length == 4).First());
    map.Add(7, samples.Where(o => o.Length == 3).First());
    map.Add(8, samples.Where(o => o.Length == 7).First());

    // Map characters with a unique appearance rate.
    segmentMap.Add('b', appearanceRates.Where(kvp => kvp.Value == 6).First().Key);
    segmentMap.Add('e', appearanceRates.Where(kvp => kvp.Value == 4).First().Key);
    segmentMap.Add('f', appearanceRates.Where(kvp => kvp.Value == 9).First().Key);

    // a is in 7 (acf) but not 1 (cf), which also exposes c.
    segmentMap.Add('a', map[7].Replace(map[1][0], '\0').Replace(map[1][1], '\0').Replace("\0", "").First());
    segmentMap.Add('c', map[7].Where(c => c != segmentMap['a'] && c != segmentMap['f']).First());

    map.Add(2, samples.Where(o => o.Length == 5 && o.Contains(segmentMap['e'])).First());
    map.Add(3, samples.Where(o => o.Length == 5 && !o.Contains(segmentMap['b']) && !o.Contains(segmentMap['e'])).First());
    map.Add(5, samples.Where(o => o.Length == 5 && o.Contains(segmentMap['b'])).First());
    map.Add(9, samples.Where(o => o.Length == 6 && !o.Contains(segmentMap['e'])).First());
    map.Add(0, samples.Where(o => o.Length == 6 && o != map[9] && o.Contains(segmentMap['c'])).First());
    map.Add(6, samples.Where(o => o.Length == 6 && o != map[9] && o != map[0]).First());


    var reverseMap = map.ToDictionary(k => Sort(k.Value), v => v.Key);
    var temp = int.Parse(
        string.Format
        (
            "{0}{1}{2}{3}",
            reverseMap[Sort(outputs[0])],
            reverseMap[Sort(outputs[1])],
            reverseMap[Sort(outputs[2])],
            reverseMap[Sort(outputs[3])]
        )
    );


    starTwoValue += temp;
    Console.WriteLine($"{temp} => {elements[1]}");
}

Console.WriteLine($"\nStar one | Appearance rate of 1, 4, 7 and 8s = {starOneAppearanceRate}");
Console.WriteLine($"Star two | Total value = {starTwoValue}");


static string Sort(string original)
{
    if (string.IsNullOrWhiteSpace(original))
        throw new Exception("Cannot sort null or empty string");

    return string.Join('\0', original.ToCharArray().OrderBy(c => c)).Replace("\0", "");
}
