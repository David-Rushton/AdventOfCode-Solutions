namespace AoC.Tokeniser
{
    public record RuleToken(
        string FieldName,
        int FirstRangeLowerBound,
        int FirstRangeUpperBound,
        int SecondRangeLowerBound,
        int SecondRangeUpperBound
    );
}
