<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateAlgorithmTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('algorithm', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name', 30)->unique()->comment('名称');
            $table->integer('created_at_ts')->index()->comment('创建时间戳');
            $table->integer('updated_at_ts')->index()->comment('更新时间戳');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('algorithm');
    }
}
