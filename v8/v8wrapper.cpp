#include <v8.h>
#include <string.h>

#include "v8wrapper.h"

using namespace v8;

char * runv8(const char *jssrc)
{
    // Get the default Isolate created at startup.
    Isolate* isolate = Isolate::GetCurrent();

    // Create a stack-allocated handle scope.
    HandleScope handle_scope(isolate);

    // Create a new context.
    Handle<Context> context = Context::New(isolate);

    // Enter the context for compiling and running the hello world script.
    Context::Scope context_scope(context);

    // Create a string containing the JavaScript source code.
    Handle<String> source = String::New(jssrc);

    // Compile the source code.
    Handle<Script> script = Script::Compile(source);

    // Run the script to get the result.
    Handle<Value> result = script->Run();

    // The JSON.stringify function object
    Handle<Object> global = context->Global();
    Handle<Object> JSON = global->Get(String::New("JSON"))->ToObject();
    Handle<Function> JSON_stringify = Handle<Function>::Cast(JSON->Get(String::New("stringify")));

    Handle<Value> args[] = { result };
    // stringify result
    Local<Value> json = JSON_stringify->Call(JSON, 1, args);

    // Convert the result to an UTF8 string and print it.
    String::Utf8Value utf8(json);

    // return result as string, must be deallocated in cgo wrapper
    return strdup(*utf8);
}
