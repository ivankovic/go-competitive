##### Signed by https://keybase.io/ivankovic
```
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2.0.22 (GNU/Linux)

iQIcBAABAgAGBQJU+3tiAAoJEHL+w/ZMh983LSQP/05J0KOhPjdMgEv1VKmz2laU
wQcuM7MtYQGkcugAMJZrSh5q0sVLinLrDiYyoPK9002Av8yLY5b8WMQPFPhls+Wo
IqBxHaHlw+HLkUVO9RB0w2TI8eBTplZr/77SeME7C7ry0L9F1spfe/VMFeO8558a
gHZ9V20mkIIljyfsPTK8xE0UKilqfEi88JunJ23qS4SiRVdJ3UdGB4ORRbkAQ3Aw
WlHriTOuUy56nru9H5EySZKU8ZGFERGVCnZhgNVwQsEvyrWG7LgGSFNJaksqq2ZX
5v7Y6vG5tJZeojrkjxiNZ3kpaK6IyvKW501GQa71SJ1m7okdhJDeCqJP8noRBoqr
AlrrIxcmGY7ghns0rSl59zuZnWXZLfvPbZ8H29DHCS13ZD2+/FM57nZ/m0zIYT/c
wTwdOSr9BkviyQihFbuNH97w+bP0ECxPn0zOe9iqRnvwyp9qqnNBTh3d/fyOoyEj
CdDsdKmoX7gPsi0obrYpezaBSUls9Lfsmr8efFy3/PgeD4WogBFedAaPHx/tl279
WOIYHMfArSdSST8SkvhzKTHWShAbGYjjyY9oWT5CYhmrGcA0yCPET/0OjcAEv67l
dk4m3DhuwRfcEemqVm+52OkGr3RS5lQKGsXuLDEZwyigh2kXVf9/TDDIhiiuDJAu
4W/7Gin+Cr3OSwsu1xmB
=Wkm+
-----END PGP SIGNATURE-----

```

<!-- END SIGNATURES -->

### Begin signed statement 

#### Expect

```
size   exec  file                     contents                                                        
             ./                                                                                       
266            .gitignore             cb568f716e3315bcebfc75bbc274b56577c2734ec3da6729ffac15919f416240
11313          LICENSE                035b55ed171df5014d81df43f4374d402ad7d7a009a47a34519856ee9c55e832
983            README.md              bb628fde9cfce7120fff707440d4ac71c4bfe02f80a46b034f79ba6423416b0a
               factorial/                                                                             
722              factorial.go         a04c8205aab2edbe82c1e95f36727d474a2a93f70e02a5b1d093eb134e9d6ce9
823              factorial_test.go    98f9b9dbf52a8a8e5b42c84e0a38b3a319930dfac955b4cca6a97b6ebba8f1fe
               gcd/                                                                                   
858              gcd.go               88ca78e6f166cd2ac2bbad77273e723f6837a2bec68dfa3b8a545bd6cdd52a30
1061             gcd_test.go          c109a0292617bca2a06435100205071da25af29a8e14e3fbef0aab63a272a8b6
               permutation/                                                                           
1190             permutation.go       ed9c0913dfc8b39e894430409eb338ecd9415ae98debb386bce837df9631faff
1794             permutation_test.go  df18e69e877d8234a36e02706040e418d6ef57215157b63098af07b6a3aa5dfe
               pow/                                                                                   
229              pow.go               8e576b047e80e9cf19bdbc7aca3890b1f19f177f03db8999502bd104694d8e9e
465              pow_test.go          2b8ccb7295e599348bcc27e760e59547f3f89d1f222426d67c1f0f8225e453b9
               prime/                                                                                 
1378             prime.go             5c05062578fa33ffa5c92bb1f58b790187f87634da2a05f63e9d2f436b4eff01
1759             prime_test.go        4c50e6f3f7aa4d7af63e5742d77277f2367646cf0c908c9935946b965b8af1c0
               swap/                                                                                  
663              swap.go              cae45f600c6deb605a88e0593da0f964f47e9cbb9e0ddf4a112012f45f172403
973              swap_test.go         5273f8fe041b94e93e3c248d2585d973bd0b92cb8f6755ed61460cd79355cfba
```

#### Ignore

```
/SIGNED.md
```

#### Presets

```
git      # ignore .git and anything as described by .gitignore files
dropbox  # ignore .dropbox-cache and other Dropbox-related files    
kb       # ignore anything as described by .kbignore files          
```

<!-- summarize version = 0.0.9 -->

### End signed statement

<hr>

#### Notes

With keybase you can sign any directory's contents, whether it's a git repo,
source code distribution, or a personal documents folder. It aims to replace the drudgery of:

  1. comparing a zipped file to a detached statement
  2. downloading a public key
  3. confirming it is in fact the author's by reviewing public statements they've made, using it

All in one simple command:

```bash
keybase dir verify
```

There are lots of options, including assertions for automating your checks.

For more info, check out https://keybase.io/docs/command_line/code_signing