using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public record TileEdge
    (
        int Top,
        int Left,
        int Bottom,
        int Right
    );


    public record Tile
    (
        int Id,
        TileEdge Edge,
        TileEdge FlippedEdge
    );
}
