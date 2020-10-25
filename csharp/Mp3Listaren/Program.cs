using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using Serilog;
using Serilog.Core;
using File = TagLib.File;

namespace Mp3Listaren
{
    static class Program
    {
        private static readonly Dictionary<string, string> SupportedVideo;

        private static readonly Dictionary<string, string> SupportedImage;

        private static readonly Dictionary<string, string> SupportedAudio;
        private static readonly Logger Log = new LoggerConfiguration()
            .WriteTo.File("log.txt")
            .CreateLogger();

        static Program()
        {
            SupportedVideo = new[]
                {
                    "mkv", "ogv", "avi", "wmv", "asf", "m4p", "m4v", "mpeg", "mpe", "mpv", "mpg", "m2v","mp4"
                }
                .Distinct()
                .ToDictionary(x => x);
            SupportedAudio = new[]
            {
                "aa", "aax", "aac", "aiff", "ape", "dsf", "flac", "m4a", "m4b", "m4p", "mp3", "mpc", "mpp", "ogg",
                "oga", "wav", "wma", "wv", "webm"
            }
                .Distinct()
                .ToDictionary(x => x);
            
            SupportedImage = new[]
                {
                    "bmp", "gif", "jpeg", "pbm", "pgm", "ppm", "pnm", "pcx", "png", "tiff", "dng", "svg"
                }
                .Distinct()
                .ToDictionary(x => x);
                
            
        }

        static void Main()
        {
            Console.WriteLine("Analyserar filer!");
            
            var fileInfos = Directory
                .GetFiles(Environment.CurrentDirectory, "*", SearchOption.AllDirectories)
                .Select(x => new FileInfo(x))
                .ToList();

            var fileInfosCount = fileInfos.Count;

            var outputPath = $@"{Directory.GetCurrentDirectory()}\filer.csv";
          
            using StreamWriter streamWriter = new StreamWriter(outputPath,false, Encoding.GetEncoding("ISO-8859-1"));
            for (var index = 0; index < fileInfos.Count; index++)
            {
                Console.Write($"\rAnalyserar fil {index+1}/{fileInfosCount}");
                
                FileInfo fileInfo = fileInfos[index];
                string extension = Path.GetExtension(fileInfo.FullName);
                var str = fileInfo.Name[..^extension.Length];

                var (duration, width, height, resolution, imageDate) = GetData(fileInfo);

                streamWriter.WriteLine($"{extension}\t" +
                                       $"{str}\t" +
                                       $"{fileInfo.Directory.FullName}\t" +
                                       $"{fileInfo.Length}\t" +
                                       $"{duration}\t",
                    $"{width}\t",
                    $"{height}\t",
                    $"{resolution}\t",
                    $"{imageDate}");
            }
        }

        private static (string Duration, string Width, string Height, string Resolution, string ImageDate) GetData(FileInfo fileInfo)
        {
            var fileType = GetFileType(fileInfo);

            if (fileType == FileType.Other)
            {
                return ("---", "---", "---", "---", "---");
            }

            try
            {
                using var file = File.Create(fileInfo.FullName);

                return fileType switch
                {
                    FileType.Image => (
                        "---",
                        file.Properties.PhotoWidth.ToString(),
                        file.Properties.PhotoHeight.ToString(),
                        $"{file.Properties.PhotoWidth}x{file.Properties.PhotoHeight}",
                        "---"),
                    FileType.Audio => (file.Properties.Duration.ToString("hh\\:mm\\:ss"), "---", "---", "---", "---"),
                    FileType.Video => (
                        file.Properties.Duration.ToString("hh\\:mm\\:ss"),
                        file.Properties.VideoWidth.ToString(),
                        file.Properties.PhotoHeight.ToString(),
                        $"{file.Properties.VideoWidth}x{file.Properties.PhotoHeight}",
                        "---"),
                    _ => throw new ArgumentOutOfRangeException()
                };
            }
            catch (Exception ex)
            {
                Log.Error(ex, "File {File} throw exception", fileInfo.FullName);
                return ("---", "---", "---", "---", "---");
            }
        }

        private static FileType GetFileType(FileInfo fileInfo)
        {
            var extension = fileInfo.Extension.TrimStart('.').ToLower();

            if (SupportedAudio.ContainsKey(extension))
            {
                return FileType.Audio;
            }
            
            if (SupportedVideo.ContainsKey(extension))
            {
                return FileType.Video;
            }
            
            if (SupportedImage.ContainsKey(extension))
            {
                return FileType.Image;
            }

            return FileType.Other;
        }
        
    }

    internal enum FileType
    {
        Image,
        Audio,
        Video,
        Other
    }
}