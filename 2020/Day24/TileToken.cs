using System;
using System.Collections.Generic;


namespace AoC
{
    public record TileToken
    (
        Colour Colour,

        Location Location,

        List<Direction> Directions
    )
    {
        public TileToken Flip() =>
            this with { Colour = this.Colour == Colour.Black ? Colour.White : Colour.Black }
        ;

        public override string ToString() =>
            string.Format
            (
                "TileToken {{ Colour = {0}, {{ {1} }} Directions = {2} }}",
                Colour,
                Location,
                string.Join(',', Directions)
            )
        ;
    }
}
