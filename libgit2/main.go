package main

/*
#cgo LDFLAGS: -lgit2
#include <stdio.h>
#include <stdlib.h>
#include <git2.h>

int TreeWalker(const git_tree_entry *entry);

int walk_cb(const char *root,
            const git_tree_entry *entry,
            void *payload)
{

	TreeWalker(entry);
	return 0;
}

int tree_walk(const git_tree *tree)
{
	int error = git_tree_walk(tree, GIT_TREEWALK_PRE, walk_cb, NULL);
	return error;
}
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"unsafe"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("repository path is required")
	}
	repo_dir := os.Args[1]
	if repo_dir == "" {
		log.Fatal("repository path is required")
	}
	crepo_dir := C.CString(repo_dir)
	defer C.free(unsafe.Pointer(crepo_dir))

	fmt.Printf("open git repository %s\n", repo_dir)
	var repo *C.git_repository
	err := C.git_repository_open(&repo, crepo_dir)
	if err != 0 {
		log.Fatalf("Failed to open git repository | path : %s status : %d", repo_dir, err)
	}

	var obj *C.git_object
	chead := C.CString("HEAD^{tree}")
	defer C.free(unsafe.Pointer(chead))

	err = C.git_revparse_single(&obj, repo, chead)
	if err != 0 {
		log.Fatalf("Failed to revparse | path : %s status : %d", repo_dir, err)
	}

	tree := (*C.git_tree)(obj)

	err = C.tree_walk(tree)
	if err != 0 {
		log.Fatalf("Failed to walk tree | path : %s status : %d", repo_dir, err)
	}

	fmt.Println("finished")
}
