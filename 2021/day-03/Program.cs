var path = args.Contains("-t") ? "Input.test.txt" : "Input.txt";

Dictionary<int, int> valueByPosition = new();
var lines = File.ReadAllLines(path);

foreach (var line in lines)
{
    for (var i = 0; i < line.Length; i++)
    {
        if (! valueByPosition.ContainsKey(i))
            valueByPosition.Add(i, 0);

        if (line[i] == '1')
            valueByPosition[i]++;
    }
}


var gamma = "";
var epsilon = "";
var threshold = lines.Length / 2;

foreach (var kvp in valueByPosition)
{
    gamma += kvp.Value > threshold ? "1" : "0";
    epsilon += gamma[kvp.Key] == '0' ? "1" : "0";
}
var powerConsumption = Convert.ToInt32(gamma, fromBase: 2) * Convert.ToInt32(epsilon, fromBase: 2);
Console.WriteLine($"Star one: Gamma {gamma}, Epsilon {epsilon}, Power consumption {powerConsumption}");



// Star two ------------------------------------------------------------------------------------------------------------


var oxygenPattern = $"{gamma[0]}";
var oxygen = "";
var co2Pattern = $"{epsilon[0]}";
var co2 = "";

for (var i = 1; i < lines[0].Length; i++)
{
    var oxygenCounter = 0;
    var oxygenFound = 0;
    var co2Counter = 0;
    var co2Found = 0;

    foreach (var line in lines)
    {
        if (oxygen == "" && line.StartsWith(oxygenPattern))
        {
            oxygenFound++;
            if (line[i] == '1')
                oxygenCounter++;
        }

        if (co2 == "" && line.StartsWith(co2Pattern))
        {
            co2Found++;
            if (line[i] == '1')
                co2Counter++;
        }
    }

    oxygenPattern += oxygenCounter >= (oxygenFound - oxygenCounter) ? '1' : '0';
    var oxygens = lines.Where(l => l.StartsWith(oxygenPattern));
    if (oxygens.Count() == 1)
        oxygen = oxygens.First();


    co2Pattern += co2Counter >= (co2Found - co2Counter) ? '0' : '1';
    var co2s = lines.Where(l => l.StartsWith(co2Pattern));
    if (co2s.Count() == 1)
        co2 = co2s.First();



    if (oxygen != "" && co2 != "")
        break;
}


Console.WriteLine($"Star two: oxygen = {oxygen} co2 = {co2} {Convert.ToUInt32(oxygen, 2) * Convert.ToUInt32(co2, 2)}");
