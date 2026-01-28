#include "binary_search_tree.h"
#include <stdlib.h>
#include <stddef.h>

int search(node_t* root, int target);
node_t *insert_node(node_t* root, int value);
size_t count_nodes(node_t *root);
void in_order_traversal(node_t *root, int *result, size_t *index);

node_t *build_tree(int *tree_data, size_t tree_data_len) {
    node_t *root = NULL;
    for (size_t i = 0; i < tree_data_len; i++) {
        root = insert_node(root, tree_data[i]);
    }
    return root;
}

void free_tree(node_t *tree) {
    if (tree == NULL) {
        return;
    }
    free_tree(tree->left);
    free_tree(tree->right);
    free(tree);
}

size_t count_nodes(node_t *root) {
    if (root == NULL) {
        return 0;
    }
    return 1 + count_nodes(root->left) + count_nodes(root->right);
}

void in_order_traversal(node_t *root, int *result, size_t *index) {
    if (root == NULL) {
        return;
    }
    
    in_order_traversal(root->left, result, index);  // Левое поддерево
    result[(*index)++] = root->data;                // Текущий узел
    in_order_traversal(root->right, result, index); // Правое поддерево
}

int *sorted_data(node_t *tree) {
    // Подсчитываем количество узлов
    size_t count = count_nodes(tree);
    
    // Выделяем память под результат
    int *result = (int*)malloc(count * sizeof(int));
    if (!result) {
        return NULL;  // Ошибка выделения памяти
    }
    
    // Заполняем массив через in-order обход
    size_t index = 0;
    in_order_traversal(tree, result, &index);
    
    return result;  // Возвращаем отсортированный массив
}

node_t *insert_node(node_t* root, int value) {
    if (root == NULL) {
        node_t* new_node = (node_t*)malloc(sizeof(node_t));
        new_node->data = value;
        new_node->left = NULL;
        new_node->right = NULL;
        return new_node;
    }
    
    if (value <= root->data) {
        root->left = insert_node(root->left, value);
    } else {
        root->right = insert_node(root->right, value);
    }
    
    return root;
}

int search(node_t* root, int target) {
    if (root == NULL) {
        return 0;
    }
    
    if (root->data == target) {
        return 1;
    }
    
    if (target < root->data) {
        return search(root->left, target);
    } else {
        return search(root->right, target);
    }
}

