using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using File = TagLib.File;

namespace Mp3Listaren
{
    class Program
    {
        private static readonly Dictionary<string, string> SupportedVideo;

        private static readonly Dictionary<string, string> SupportedImage;

        private static readonly Dictionary<string, string> SupportedAudio;
        
        static Program()
        {
            SupportedVideo = new[]
                {
                    "mkv", "ogv", "avi", "wmv", "asf", "m4p", "m4v", "mpeg", "mpe", "mpv", "mpg", "m2v",
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
            var fileInfos = Directory
                .GetFiles(Environment.CurrentDirectory, "*", SearchOption.AllDirectories)
                .Select(x => new FileInfo(x));

            var outputPath = $@"{Directory.GetCurrentDirectory()}\filer.txt";
          
            using StreamWriter streamWriter = new StreamWriter(outputPath);
            foreach (FileInfo fileInfo in fileInfos)
            {
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

        private static FileType GetFileType(FileInfo fileInfo)
        {
            var extension = fileInfo.Extension.TrimStart('.');

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