package main

/*================================= imports ======================*/
import (
  "fmt"
  "os"
  "io"
  "strconv"
  "os/exec"
  "bufio"
)

var progname string
const MAX_INT int = 1 << 32 - 1
const LINE_SIZE int = 1024
const INBUFSIZ int = 16 * 1024

/*================================= types =========================*/
type selpg_args struct {
  start_page int
  end_page int
  in_filename string
  page_len int  /* default value, can be overriden by "-l number" on command line */
  page_type int  /* 'l' for lines-delimited, 'f' for form-feed-delimited */
      /* default is 'l' */
  print_dest string
}

// get the name of the program
func get_Progname(name string) string {
  var pos = 0
  for i, ch := range name {
    if ch == '/' { pos = i }
  }
  return name[pos:]
}

// main()function
func main()  {
  /* save name by which program is invoked, for error messages */
  progname = get_Progname(os.Args[0])

  // initialize
  var sa selpg_args
  sa.start_page = -1
  sa.end_page = -1
  sa.in_filename = ""
  sa.page_len = 72
  sa.page_type = 'l'
  sa.print_dest = ""

  process_args(&sa)
  process_input(sa)
}

/*================================= process_args() ================*/
// set sa
func process_args(sa *selpg_args) {
  var s string
  var i int

  // error 1: not enough arguments
  if len(os.Args) < 3 {
    fmt.Fprintf(os.Stderr, "%s: not enough arguments\n",
    progname)
    help()
    os.Exit(1);
  }

  // first argument
  s = os.Args[1]
  // error 2: wrong format
  if len(s) < 2 || s[:2] != "-s" {
    fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstart_page\n",
    progname)
    help()
    os.Exit(2);
  }
  i, _ = strconv.Atoi(s[2:])
  // error 3: overflow
  if i < 1 || i > MAX_INT {
    fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n",
    progname, s[2:])
    help()
    os.Exit(3);
  }
  sa.start_page = i

  // second argument
  s = os.Args[2]
  // error 4: wrong format
  if len(s) < 2 || s[:2] != "-e" {
    fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n",
    progname)
    help()
    os.Exit(4);
  }
  i, _ = strconv.Atoi(s[2:])
  // error 5: overflow
  if i < 1 || i > MAX_INT {
    fmt.Fprintf(os.Stderr, "%s: invalid end page %s\n",
    progname, s[2:])
    help()
    os.Exit(5);
  }
  sa.end_page = i

  argnum := 3;
  // option
  for _, s = range os.Args[3:] {
    if s[0] != '-' { break }
    argnum++
    switch s[1] {
    case 'l':
      i, _ = strconv.Atoi(s[2:])
      // error 6: 数字超过int范围
      if i < 1 || i > MAX_INT {
        fmt.Fprintf(os.Stderr, "%s: invalid page length %s\n",
        progname, s[2:])
        help()
        os.Exit(6);
      }
      sa.page_len = i
    case 'f':
      // error 7: option should be "-f"
      if s != "-f" {
        fmt.Fprintf(os.Stderr, "%s: option should be \"-f\"\n",
        progname)
        help()
        os.Exit(7);
      }
      sa.page_type = 'f'
    case 'd':
      // error 8: -d option requires a printer destination
      if s == "-d" {
        fmt.Fprintf(os.Stderr, "%s: -d option requires a printer destination\n",
        progname)
        help()
        os.Exit(8);
      }
      sa.print_dest = s[2:]
    default:
      // error 9: unknown option
      fmt.Fprintf(os.Stderr, "%s: unknown option %s\n",
      progname, s)
      help()
      os.Exit(9);
    }
  }

  // input file
  if argnum < len(os.Args) {
    s = os.Args[argnum]
    sa.in_filename = s
    // error 10: input file does not exist
    if _, err := os.Stat(s); err != nil && os.IsNotExist(err) {
      fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n",
      progname, s)
      os.Exit(10);
    }
  }
}

/*================================= process_input() ===============*/
func process_input(sa selpg_args) {
  var err error
  var s string
  var page int
  var line int
  var r *bufio.Reader
  var w *bufio.Writer
  var rw *bufio.ReadWriter
  var cmd *exec.Cmd
  var file *os.File

  // if not standard
  if sa.in_filename != "" {
    file, err = os.OpenFile(sa.in_filename, os.O_RDONLY, 0)
    r = bufio.NewReaderSize(file, INBUFSIZ)
    // error 11: can not open the file
    if os.IsNotExist(err) {
      fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n",
			progname, sa.in_filename)
      os.Exit(11);
    }
  } else {
    r = bufio.NewReaderSize(os.Stdin, INBUFSIZ)
  }

  // if not standart
  if sa.print_dest != "" {
    // run lp -d
    cmd = exec.Command("lp", "-d" + sa.print_dest)
    fout, err := cmd.StdinPipe()
    w = bufio.NewWriterSize(fout, INBUFSIZ)
    // error 12: can not open pipe
    if err != nil {
      fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n",
      progname, "ls " + s)
      os.Exit(12);
    }
  } else {
    w = bufio.NewWriterSize(os.Stdout, INBUFSIZ)
  }

  // combine r and w
  rw = bufio.NewReadWriter(r, w)

  // start running
  // set pattern
  if sa.page_type == 'l' {
    line = 0
    page = 1

    for true {
      // used for bytes，1024 bytes for each line
      crc := make([]byte, LINE_SIZE)
      _, err = rw.Read(crc)

      // error : input stream error
      if err != nil {
        if err == io.EOF {
          break
        } else {
          fmt.Fprintf(os.Stderr,
            "%s: input stream error\n", progname)
        }
      }
      line++
      if (line > sa.page_len) {
        page++
        line = 1
      }
      // page number
      if page >= sa.start_page && page <= sa.end_page {
        rw.WriteString(string(crc))
      }
    }
  } else {
    page = 1

    for true {
      // one byte one time
      ch, _, err := rw.ReadRune()
      // error : imput stream error
      if err != nil {
        if err == io.EOF {
          break
        } else {
          fmt.Fprintf(os.Stderr,
            "%s: input stream error\n", progname)
        }
      }
      if ch == '\f' { page++ }
      // page number
      if page >= sa.start_page && page <= sa.end_page {
        rw.WriteRune(ch)
      }
    }
  }
  // flush
  rw.Flush()

  // end the output
  if page < sa.start_page {
    fmt.Fprintf(os.Stderr,
      "%s: start_page (%d) greater than total pages (%d)," +
  		" no output written\n", progname, sa.start_page, page)
  } else if page < sa.end_page {
    fmt.Fprintf(os.Stderr,
      "%s: end_page (%d) greater than total pages (%d)," +
  		" less output than expected\n", progname, sa.end_page, page)
  } else {
    // if input files
    if sa.in_filename != "" { file.Close() }
    if sa.print_dest != "" {
      out, _ := cmd.CombinedOutput()
      fmt.Fprint(os.Stderr, string(out))
    }
    // over
    fmt.Fprintf(os.Stderr, "\n%s: done\n", progname)
  }
}

/*================================= help() =======================*/
// tips
func help()  {
  fmt.Fprintf(os.Stderr,"\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ]" +
	" [ -ddest ] [ in_filename ]\n", progname)
}

/*================================= EOF ===========================*/
