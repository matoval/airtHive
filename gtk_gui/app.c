#include <gtk/gtk.h>
#include <stdio.h>

static void activate (GtkApplication *app, gpointer        user_data) {
  GtkWidget *window = gtk_application_window_new(app);
  GtkWidget *grid = gtk_grid_new();

  gtk_window_set_title (GTK_WINDOW (window), "airtHive");
  gtk_window_set_default_size (GTK_WINDOW (window), 800, 600);

  gtk_grid_set_column_spacing(GTK_GRID(grid), 10);
  gtk_grid_get_row_spacing(GTK_GRID(grid));

  GtkWidget *button1 = gtk_button_new_with_label ("Chat");
  GtkWidget *button2 = gtk_button_new_with_label ("Backends");

  gtk_grid_attach(GTK_GRID(grid), button1, 0,0,1,1);
  gtk_grid_attach(GTK_GRID(grid), button2, 1,0,1,1);

  gtk_window_set_child(GTK_WINDOW(window), grid);
 
  gtk_window_present(GTK_WINDOW(window));
}

int main(int argc, char **argv) {
  GtkApplication *app;
  int status;

  app = gtk_application_new ("org.example.airtHive", G_APPLICATION_DEFAULT_FLAGS);
  g_signal_connect (app, "activate", G_CALLBACK (activate), NULL);
  status = g_application_run (G_APPLICATION (app), argc, argv);
  g_object_unref (app);

  return status;
}
