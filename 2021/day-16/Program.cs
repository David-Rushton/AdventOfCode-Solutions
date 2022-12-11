const int MinimumPacketLength = 11;
// var testInput = args[1];
// var testInput = "D2FE28";
// var testInput = "38006F45291200";
var testInput = "A0016C880162017C3686B18A3D4780";
var input = "220D700071F39F9C6BC92D4A6713C737B3E98783004AC0169B4B99F93CFC31AC4D8A4BB89E9D654D216B80131DC0050B20043E27C1F83240086C468A311CC0188DB0BA12B00719221D3F7AF776DC5DE635094A7D2370082795A52911791ECB7EDA9CFD634BDED14030047C01498EE203931BF7256189A593005E116802D34673999A3A805126EB2B5BEEBB823CB561E9F2165492CE00E6918C011926CA005465B0BB2D85D700B675DA72DD7E9DBE377D62B27698F0D4BAD100735276B4B93C0FF002FF359F3BCFF0DC802ACC002CE3546B92FCB7590C380210523E180233FD21D0040001098ED076108002110960D45F988EB14D9D9802F232A32E802F2FDBEBA7D3B3B7FB06320132B0037700043224C5D8F2000844558C704A6FEAA800D2CFE27B921CA872003A90C6214D62DA8AA9009CF600B8803B10E144741006A1C47F85D29DCF7C9C40132680213037284B3D488640A1008A314BC3D86D9AB6492637D331003E79300012F9BDE8560F1009B32B09EC7FC0151006A0EC6082A0008744287511CC0269810987789132AC600BD802C00087C1D88D05C001088BF1BE284D298005FB1366B353798689D8A84D5194C017D005647181A931895D588E7736C6A5008200F0B802909F97B35897CFCBD9AC4A26DD880259A0037E49861F4E4349A6005CFAD180333E95281338A930EA400824981CC8A2804523AA6F5B3691CF5425B05B3D9AF8DD400F9EDA1100789800D2CBD30E32F4C3ACF52F9FF64326009D802733197392438BF22C52D5AD2D8524034E800C8B202F604008602A6CC00940256C008A9601FF8400D100240062F50038400970034003CE600C70C00F600760C00B98C563FB37CE4BD1BFA769839802F400F8C9CA79429B96E0A93FAE4A5F32201428401A8F508A1B0002131723B43400043618C2089E40143CBA748B3CE01C893C8904F4E1B2D300527AB63DA0091253929E42A53929E420";
var message = args.Contains("--test") ? testInput : input;
var packets = new Queue<char>(message.Select(c => c));
var bits = new Queue<char>();


// convert packets to bits
while (packets.Any())
{
    foreach (var bit in ConvertHexToBinaryChar(packets.Dequeue()))
    {
        bits.Enqueue(bit);
    }
}


// read bits
ReadPackets(bits);


static void ReadPackets(Queue<char> bits)
{
    long versionSum = 0;
    long valueSum = 0;

    while (bits.Count >= MinimumPacketLength)
    {
        (long versions, long value) = ReadPacket(bits);
        valueSum += value;
        versionSum += versions;
    }

    Console.WriteLine($"Version: {versionSum} | Value: {valueSum}");
}

static (long versions, long value) ReadPacket(Queue<char> bits, int level = 0)
{
    var version = ConvertBinaryToInt(TakeFromQueue(bits, 3));
    var typeId = ConvertBinaryToInt(TakeFromQueue(bits, 3));
    var values = new List<long>();
    long value = 0;

    if (typeId == 4)
    {
        var binaryValue = string.Empty;
        var lastBlock = false;
        while (!lastBlock)
        {
            lastBlock = TakeFromQueue(bits, 1) == "0";
            binaryValue += TakeFromQueue(bits, 4);
        }

        value = ConvertBinaryToInt(binaryValue);

        var indent = new string(' ', level * 2);
        Console.WriteLine($"{level} {indent}Literal value: {value}");
    }
    else
    {
        var typeLengthId = ConvertBinaryToInt(TakeFromQueue(bits, 1));

        if (typeLengthId == 0)
        {
            // 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet
            var lengthOfSubPackets = ConvertBinaryToInt(TakeFromQueue(bits, 15));
            var targetLength = bits.Count - lengthOfSubPackets;

            while (bits.Count > targetLength)
            {
                (long versions, long packetValue) = ReadPacket(bits, level + 1);
                version += versions;
                values.Add(packetValue);
            }
        }
        else
        {
            // 11 bits are a number that represents the number of sub-packets immediately contained by this packet
            var countOfSubPackets = ConvertBinaryToInt(TakeFromQueue(bits, 11));
            while (countOfSubPackets > 0)
            {
                (long versions, long packetValue) = ReadPacket(bits, level + 1);
                version += versions;
                values.Add(packetValue);
                countOfSubPackets--;
            }
        }

        value = (typeId) switch
        {
            0 => values.Sum(),
            1 => values.Aggregate((acc, next) => acc * next),
            2 => values.Min(),
            3 => values.Max(),
            5 => values[0] > values[1] ? 1 : 0,
            6 => values[0] < values[1] ? 1 : 0,
            7 => values[0] == values[1] ? 1 : 0,
            _ => throw new Exception($"Type id not supported: {typeId}")
        };
    }


    return (version, value);
}

static string TakeFromQueue(Queue<char> queue, int count)
{
    var result = string.Empty;
    while (count > 0)
    {
        result += queue.Dequeue();
        count--;
    }

    return result;
}

static string ConvertHexToBinaryChar(char hex) => ConvertHexToBinary(hex.ToString());
static string ConvertHexToBinary(string hex)
{
    var result = string.Empty;
    foreach (var c in hex)
    {
        result += (c) switch
        {
            '0' => "0000",
            '1' => "0001",
            '2' => "0010",
            '3' => "0011",
            '4' => "0100",
            '5' => "0101",
            '6' => "0110",
            '7' => "0111",
            '8' => "1000",
            '9' => "1001",
            'A' => "1010",
            'B' => "1011",
            'C' => "1100",
            'D' => "1101",
            'E' => "1110",
            'F' => "1111",
            _ => throw new Exception($"Invlaid hex character {c} in {hex}")
        };

    }
    return result;
}

static long ConvertBinaryToInt(string binary)
{
    var bits = new Stack<char>(binary.Select(c => c));
    long result = 0;
    long bitValue = 1;

    while (bits.Any())
    {
        var bit = bits.Pop();

        if (bit == '1')
            result += bitValue;

        bitValue *= 2;
    }

    return result;
}
