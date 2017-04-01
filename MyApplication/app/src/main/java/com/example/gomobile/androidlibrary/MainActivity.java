package com.example.gomobile.androidlibrary;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.TextView;

import androidlibrary.Androidlibrary;

public class MainActivity extends AppCompatActivity
{
    private EditText viewEdit;
    private TextView viewText;

    @Override
    protected void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        viewEdit = (EditText) findViewById(R.id.editText);
        viewText = (TextView) findViewById(R.id.viewText);

    }

    public void talkToGo(View view)
    {
        viewText.setText(Androidlibrary.echo(viewEdit.getText().toString()));
    }
}
