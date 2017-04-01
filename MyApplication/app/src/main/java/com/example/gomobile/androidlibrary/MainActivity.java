package com.example.gomobile.androidlibrary;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.CheckedTextView;
import android.widget.EditText;
import android.widget.TextView;

import androidlibrary.Androidlibrary;
import androidlibrary.JavaCallback;

public class MainActivity extends AppCompatActivity
{
    final static String TAG = MainActivity.class.getSimpleName();

    private EditText viewEdit;
    private TextView viewText;
    private CheckedTextView checkText;
    private JavaCallback javaCallback;

    @Override
    protected void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        viewEdit = (EditText) findViewById(R.id.editText);
        viewText = (TextView) findViewById(R.id.viewText);
        checkText = (CheckedTextView) findViewById(R.id.checkedTextView);

        javaCallback = new JavaCallback()
        {
            @Override
            public void heartbeat()
            {
                runOnUiThread(new Runnable()
                {
                    @Override
                    public void run()
                    {
                        checkText.toggle();
                    }
                });
            }

            @Override
            public void relayMessage(String s, long l)
            {
                Log.d(TAG, "Got a message from Go: " + s + ", value " + l);
            }
        };
    }

    public void talkToGo(View view)
    {
        viewText.setText(Androidlibrary.echo(viewEdit.getText().toString()));
    }

    public void doRegister(View view)
    {
        Androidlibrary.registerCallback(javaCallback);
    }

    public void doUnregister(View view)
    {
        Androidlibrary.unregisterCallback();
    }
}
