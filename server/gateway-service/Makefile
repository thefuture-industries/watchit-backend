# Makefile для Windows (mingw32-make)
# Он находит все .c файлы, компилирует их в объектные файлы (сохраняя в bin),
# затем линкует объектные файлы в исполняемый файл, и позволяет запустить сборку.

# Компилятор и флаги сборки

CC = gcc
CFLAGS = -Wall -Wextra -O2
LDFLAGS = -lpthread

# Каталог для объектных файлов
OBJ_DIR = obj

# Имя исполняемого файла
TARGET = $(OBJ_DIR)/gateway-server.exe

# Получаем список всех C-файлов в текущем каталоге и подкаталогах
SOURCES := $(shell find . -name "*.c" -not -path "*/test/*" -print)

# Создаем список объектных файлов
OBJECTS := $(patsubst %.c,$(OBJ_DIR)/%.o,$(SOURCES))

# Основная цель
all: $(TARGET)

# Правило для создания исполняемого файла
$(TARGET): $(OBJECTS)
	@echo "Linking..."
	$(CC) $(OBJECTS) -lws2_32 $(LDFLAGS) -o $(TARGET)

# Правило для создания объектных файлов
$(OBJ_DIR)/%.o: %.c
	@echo "Compiling $<"
	@mkdir -p $(@D)  # Создаем каталог для объектных файлов, если его нет
	$(CC) $(CFLAGS) -c $< -o $@

# Цель для очистки сборки
clean:
	@echo "Cleaning..."
	rm -rf $(OBJ_DIR) $(TARGET)

# Запуск приложения
run: $(TARGET)
	./$(TARGET)

.PHONY: all clean run
