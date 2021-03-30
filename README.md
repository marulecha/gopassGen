# passGenerator written in golang

mainly for self development purpose.

<p>Usage: 

$ ./passGen -d 12 -s -u
  -d int
        Number of password digits. (default 8)
        
  -s    Include symbols.
  
  -u    Include upper case digits.</p>

<p>Example output:

$ ./passGen -d 20 -u -s

Digits: 20

Symbols: true

Upper case: true


Password: !N}o65f|r06Ep'@/w$LB
</p>
