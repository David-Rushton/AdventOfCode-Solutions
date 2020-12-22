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


            // format: Tile dddd:
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

            return new Tile(
                id: tileId,
                top: tileLines.First(),
                left: tileLines.Select(c => c.First()).ToArray(),
                bottom: tileLines.Last(),
                right: tileLines.Select(c => c.Last()).ToArray()
            );
        }
    }
}
