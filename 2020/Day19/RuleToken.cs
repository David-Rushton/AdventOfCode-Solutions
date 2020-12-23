namespace AoC
{
    public record SubRules(
        int[] Value
    );

    public record RuleToken
    (
        int Id,
        SubRules[] SubRules,
        string Value
    )
    {
        public bool IsRuleZero => Id is 0;

        public bool IsLeaf => Value is "a" or "b";


        public override string ToString()
        {
            var subRulesTxt = string.Join(' ', SubRules[0].Value);

            if(SubRules.Length == 2)
                subRulesTxt += $" | {string.Join(' ', SubRules[1].Value)}";

            return string.Format
            (
                "RuleToken {{ Id = {0}, SubRules = {{ {1} }}, Value = {2} }}",
                this.Id,
                subRulesTxt,
                this.Value
            );
        }
    }
}
