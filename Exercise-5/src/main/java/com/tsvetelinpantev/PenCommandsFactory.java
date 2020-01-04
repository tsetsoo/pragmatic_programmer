package com.tsvetelinpantev;

import java.util.HashMap;
import java.util.Map;

public class PenCommandsFactory {
    private final Map<String, Command> commands = new HashMap<>();

    private PenCommandsFactory() {
    }

    public static PenCommandsFactory newInstance() {
        PenCommandsFactory factory = new PenCommandsFactory();
        factory.commands.put("P", (String... parameters) -> {
            if (parameters.length != 1) {
                throw new IllegalArgumentException("Select pen must have one parameter");
            }
            return String.format("select pen %s", parameters[0]);
        });
        factory.commands.put("D", (String... parameters) -> {
            if (parameters.length != 0) {
                throw new IllegalArgumentException("Pen down must have no parameters");
            }
            return "pen down";
        });
        factory.commands.put("W", (String... parameters) -> {
            if (parameters.length != 1) {
                throw new IllegalArgumentException("Draw west must have one parameter");
            }
            return String.format("draw west %scm", parameters[0]);
        });
        factory.commands.put("E", (String... parameters) -> {
            if (parameters.length != 1) {
                throw new IllegalArgumentException("Draw east must have one parameter");
            }
            return String.format("draw east %scm", parameters[0]);
        });
        factory.commands.put("N", (String... parameters) -> {
            if (parameters.length != 1) {
                throw new IllegalArgumentException("Draw north must have one parameter");
            }
            return String.format("draw north %scm", parameters[0]);
        });
        factory.commands.put("S", (String... parameters) -> {
            if (parameters.length != 1) {
                throw new IllegalArgumentException("Draw south must have one parameter");
            }
            return String.format("draw south %scm", parameters[0]);
        });
        factory.commands.put("U", (String... parameters) -> {
            if (parameters.length != 0) {
                throw new IllegalArgumentException("Pen down must have no parameters");
            }
            return "pen up";
        });
        return factory;
    }

    public String execute(String commandName, String... parameters) {
        if (!commands.containsKey(commandName)) {
            throw new IllegalArgumentException("No such command " + commandName + ". Available commands are: " + commands.keySet());
        }
        return commands.get(commandName).apply(parameters);
    }
}
