using System.Collections.Generic;


namespace AoC
{
    public record Food
    (
        List<string> Ingredients,
        List<string> Allergens
    )
    {
        public override string ToString() =>
            string.Format
                (
                    "Food {{ Ingredients {{ {0} }} Allergens {{ {1} }} }}",
                    string.Join(", ", Ingredients),
                    string.Join(", ", Allergens)
                )
            ;
    }
}
