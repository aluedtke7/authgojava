package de.test;

import org.springframework.core.env.SimpleCommandLinePropertySource;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

public class App {
    public static void main(String[] args) {
        SimpleCommandLinePropertySource clps = new SimpleCommandLinePropertySource(args);
        if (!clps.containsProperty("password")) {
            System.out.println("Please supply a password via --password");
            System.out.println("Or compare a password and a hashed password via --hashedpw");
            System.out.println("You should use single quotes (') for the hashed password");
            System.exit(1);
        }

        BCryptPasswordEncoder passwordEncoder = new BCryptPasswordEncoder(12);
        String password = clps.getProperty("password");

        if (clps.containsProperty("hashedpw")) {
            System.out.println("Comparing password and hashed password:");
            String hashedpw = clps.getProperty("hashedpw");
            if (passwordEncoder.matches(password, hashedpw)) {
                System.out.println("Password and hashed password do match!");
            } else {
                System.out.println("Password and hashed password do NOT match!");
            }
        } else {
            String hashedPassword = passwordEncoder.encode(password);
            System.out.println("Original:    " + password);
            System.out.println("Salted hash: " + hashedPassword);
        }
    }
}
