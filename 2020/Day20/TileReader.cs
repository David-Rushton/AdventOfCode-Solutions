using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.IO;


namespace AoC
{
    public class TileReader
    {
        const int _invalidTileId = int.MinValue;


        public IEnumerable<Tile> GetTiles(string inputPath)
        {
            var tileLines = new List<char[]>();
            var tileId = _invalidTileId;

            foreach(var line in File.ReadLines(inputPath).Append(string.Empty))
                switch (line)
                {
                    // tile id
                    case string s when s.StartsWith("Tile "):
                        tileId = ExtractTileId(line);
                        break;

                    // end of tile
                    case string s when s == string.Empty:
                        yield return ConvertRawInputToTitle(tileId, tileLines);
                        tileLines.Clear();
                        tileId = _invalidTileId;
                        break;

                    // next row of current tile
                    default:
                        tileLines.Add(line.ToCharArray());
                        break;
                }


            Debug.Assert(tileLines.Count == 0, "Input read but not returned");


            int ExtractTileId(string line) =>
                int.Parse
                (
                    line.Split(' ')[1].Replace(':', '\0')
                );
            ;
        }

        private Tile ConvertRawInputToTitle(int tileId, List<char[]> tileLines)
        {
            Debug.Assert(tileId != _invalidTileId, "Invalid tile id");

            var Top     = tileLines.First();
            var Left    = tileLines.Select(c => c.First()).Reverse().ToArray();
            var Bottom  = tileLines.Last().Reverse().ToArray();
            var Right   = tileLines.Select(c => c.Last()).ToArray();


            // We need to reverse left and bottom edges to account for the effect of tile rotation
            return new Tile(
                Id: tileId,
                Edge: new TileEdge
                (
                    ConvertEdgeToValue(Top),
                    ConvertEdgeToValue(Left),
                    ConvertEdgeToValue(Bottom),
                    ConvertEdgeToValue(Right)
                ),
                FlippedEdge: new TileEdge
                (
                    ConvertEdgeToValue(Top.Reverse().ToArray()),
                    ConvertEdgeToValue(Left.Reverse().ToArray()),
                    ConvertEdgeToValue(Bottom.Reverse().ToArray()),
                    ConvertEdgeToValue(Right.Reverse().ToArray())
                )
            );
        }

        private int ConvertEdgeToValue(char[] edge)
        {
            var total = 0;
            var characterValue = 1;

            foreach(var character in edge)
            {
                if(character == '#')
                    total += characterValue;

                characterValue = characterValue * 2;
            }

            return total;
        }
    }
}
;
