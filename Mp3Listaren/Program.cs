using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace Mp3Listaren
{
    class Program
    {
        private static readonly Dictionary<string, string> SupportedFiles;

        static Program()
        {
            SupportedFiles = new[]
            {
                // video
                "mkv", "ogv", "avi", "wmv", "asf", "mp4", "m4p", "m4v", "mpeg", "mpg", "mpe", "mpv", "mpg", "m2v",
                // audio
                "aa", "aax", "aac", "aiff", "ape", "dsf", "flac", "m4a", "m4b", "m4p", "mp3", "mpc", "mpp", "ogg",
                "oga", "wav", "wma", "wv", "webm"
            }.ToDictionary(x => x);
        }
         
        static void Main(params string[] args)
        {
            var fileInfos = Directory
                .GetFiles(Environment.CurrentDirectory, "*", SearchOption.AllDirectories)
                .Select(x => new FileInfo(x));

            var outputPath = $@"{Directory.GetCurrentDirectory()}\filer.txt";

            if (args.Contains("--debug"))
            {
                Console.WriteLine($"Skriver till '{outputPath}'");
            }
            
            using StreamWriter streamWriter = new StreamWriter(outputPath);
            foreach (FileInfo fileInfo in fileInfos)
            {
                string extension = Path.GetExtension(fileInfo.FullName);
                var str = fileInfo.Name[..^extension.Length];
                
                streamWriter.WriteLine($"{extension}\t{str}\t{fileInfo.Directory.FullName}\t{fileInfo.Length}\t{GetLength(fileInfo)}");
            }
        }

        private static string GetLength(FileInfo fileInfo)
        {
            string extension = Path.GetExtension(fileInfo.FullName).TrimStart('.');

            if (SupportedFiles.ContainsKey(extension))
            {
                using var file = TagLib.File.Create(fileInfo.FullName);
                return file.Properties.Duration.ToString("hh\\:mm\\:ss");
            }

            return "---";
        }
    }
}