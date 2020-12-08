using System;


namespace AoC.PassportRules
{
    public class HeightPassportRule: IPassportRule
    {
        public string FieldName => "hgt";

        public bool IsValidValue(string fieldValue)
        {
            var validRanges = new []
            {
                new { unit = "cm", minHeight = 150, maxHeight = 193 },
                new { unit = "in", minHeight =  59, maxHeight =  76 },
            };

            foreach(var validRange in validRanges)
            {
                if(fieldValue.EndsWith(validRange.unit))
                {
                    if(int.TryParse(fieldValue.Remove(fieldValue.Length -2), out var heightInUnits))
                        return heightInUnits >= validRange.minHeight && heightInUnits <= validRange.maxHeight;
                }
            }

            return false;
        }
    }
}
