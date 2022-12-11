using System.Linq;

namespace Day08;

public class SevenSegmentDigit
{
    private static readonly char[] _allSegments = new[] { 'a', 'b', 'c', 'd', 'e', 'f', 'g' };
    private static Dictionary<int, SevenSegmentDigit> _segments = new();

    public SevenSegmentDigit(int number, char[] segments)
    {
        Number = number;
        Length = segments.Length;
        Segments = segments;
        DoesNotcontain = _allSegments.Where(s => !segments.Contains(s)).ToArray();
    }


    public int Number { get; private set; }
    public int Length { get; private set; }
    public char[] Segments { get; private set; }
    public char[] DoesNotcontain { get; private set; }


    public static SevenSegmentDigit[] Build()
    {
        _segments.Add(0, new SevenSegmentDigit(0, new[] { 'a', 'b', 'c',      'e', 'f', 'g' } )); // 6
        _segments.Add(1, new SevenSegmentDigit(1, new[] {           'c',           'f'      } )); // 2
        _segments.Add(2, new SevenSegmentDigit(2, new[] { 'a',      'c', 'd', 'e',      'g' } )); // 5
        _segments.Add(3, new SevenSegmentDigit(3, new[] { 'a',      'c', 'd',      'f', 'g' } )); // 5
        _segments.Add(4, new SevenSegmentDigit(4, new[] {      'b', 'c', 'd',      'f'      } )); // 4
        _segments.Add(5, new SevenSegmentDigit(5, new[] { 'a', 'b',      'd',      'f', 'g' } )); // 5
        _segments.Add(6, new SevenSegmentDigit(6, new[] { 'a', 'b',      'd', 'e', 'f', 'g' } )); // 6
        _segments.Add(7, new SevenSegmentDigit(7, new[] { 'a',      'c',           'f'      } )); // 3
        _segments.Add(8, new SevenSegmentDigit(8, new[] { 'a', 'b', 'c', 'd', 'e', 'f', 'g' } )); // 7
        _segments.Add(9, new SevenSegmentDigit(9, new[] { 'a', 'b', 'c', 'd',      'f', 'g' } )); // 6
                                                        // 2    4    2    3    6    1    3
                                                        // 8    6    8    7    4    9    7
                                                        // .    !    .    .    !    !    .

        return _segments.Select(s => s.Value).ToArray();
    }
}
