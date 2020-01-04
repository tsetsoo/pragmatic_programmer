package com.tsvetelinpantev;

@FunctionalInterface
public interface Command {
    String apply(String... parameters);
}
