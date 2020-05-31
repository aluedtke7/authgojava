# Password hashing with Java and Go
This project contains two sub projects: one for a Java (Spring) implementation
and one for a Go implementation with the standard Go crypto library.

Both programs are command line tools and you can either generate a salted, 
hashed password from a given (raw) password or you can compare the raw
password with the hashed one.

Both implementations use the bcrypt algorithm with a default strength of 12.

