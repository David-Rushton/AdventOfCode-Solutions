// using System;
// using System.Collections.Generic;
// using System.Diagnostics;
// using System.Linq;


// namespace AoC
// {
//     public record TileBoarder
//     (
//         string Top,
//         string Left,
//         string Bottom,
//         string Right
//     );


//     public class TileEx
//     {
//         static Dictionary<int, TileEx> _tiles = new();


//         public TileEx(int id, string top, string left, string, string bottom, string right)
//         {
//             (Id, Boarders, IsFlipped) = (id, new TileBoarder(top, left, bottom, right), false);
//             _tiles.Add(id, this);
//         }


//         public int Id { get; init; }

//         public TileBoarder Boarders { get; private set; }

//         public bool IsFlipped { get; private set; }


//         public void Flip()
//         {
//             Boarders = Boarders with
//             {
//                 Top = Boarders.Bottom.Reverse().ToString(),
//                 Left = Boarders.Right.Reverse().ToString(),
//                 Bottom = Boarders.Top.Reverse().ToString(),
//                 Right = Boarders.Right.Reverse().ToString()
//             };
//             IsFlipped = ! IsFlipped;
//         }

//         public void GetMatches()
//         {
//             foreach(var tile in _tiles.Where(t => t.Key != this.Id))
//             {

//             }
//         }


//         public class TileMatches
//         {
//             readonly TileEx _parent;


//             private TileMatches(TileEx parent) => _parent = parent;


//             public IEnumerable<Tile> Top => _tiles.Where(t => _parent.Boarders.Top == t.Value.Boarders.
//             public IEnumerable<Tile> Top => _tiles.Where(t => this.)
//             public IEnumerable<Tile> Top => _tiles.Where(t => this.)
//             public IEnumerable<Tile> Top => _tiles.Where(t => this.)
//         }
//     }

// }
