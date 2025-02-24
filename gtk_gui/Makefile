# Compiler and flags
CC = gcc
CFLAGS = `pkg-config --cflags --libs gtk4`
SRC = gtk_gui/app.c
TARGET = gtk_gui/app

gtk-app:
	$(CC) $(SRC) -o $(TARGET) $(CFLAGS) $(LDFLAGS)

clean:
	rm -f $(TARGET)