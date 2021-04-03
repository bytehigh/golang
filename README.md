# Picturesque Assembler converter

A very simple conversion program that converts ZX Spectrum Picturesque Assembler files into ASCII. 

Picturesque assembler files are almost completely ASCII already, only the line numbers need conversion.

Command line arguments:
-f input_filename
-o output_filename

<ul>
  <li>If no input file is provided the program will terminate with an open file error.</li>
  <li>If no output file is provided the program will write to "output.asm" by default.</li>
</ul>

If the output file already exists it will be overwritten, so please be careful!

File format:
<ul>
  <li>File header of 128 bytes</li>
  <li>Next byte is 0xOD - each line that follows is terminated by 0xOD</li>
  <li>First two bytes of each line hold a representation of the line number in hex</li>
  <li>EOF denoted by 0xFF</li>
</ul>
  
Examples of two byte line numbers in hex:
<ul>
  <li>0010 - line 10</li>
  <li>0040 - line 40</li>
  <li>0120 - line 120</li>
</ul>
and so on...
<br>
The conversion programme throws away the file header and then reads the rest of the file one byte at a time.
The first two bytes of each line (following a carriage return 0x0D) are converted to ASCII
Everything else is written out as-is.

