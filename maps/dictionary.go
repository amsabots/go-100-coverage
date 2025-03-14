package maps

import "errors"


var ErrKeyNotFoundInMap = errors.New("word not found")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	defination, ok := d[word];
	if !ok {return "", ErrKeyNotFoundInMap}
	return defination, nil
}


func (d Dictionary) Add(word, defination string){
	d[word] = defination
}